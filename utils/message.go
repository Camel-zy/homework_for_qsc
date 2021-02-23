package utils

import (
	"errors"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetMessageBalance() (float32, error) {
	//TODO(TO/GA): call sms service
	var cost float32

	return cost, nil
}

func GenerateMap(templateMap *map[string]string, answer *model.Answer, departmentID uint, interviewID uint) error {
	// answer
	{
		if _, ok := (*templateMap)["#name#"]; ok {
			(*templateMap)["#name#"] = answer.Name
		}

		if _, ok := (*templateMap)["#stuid#"]; ok {
			(*templateMap)["#stuid#"] = answer.ZJUid
		}

		if _, ok := (*templateMap)["#phone#"]; ok {
			(*templateMap)["#phone#"] = answer.Mobile
		}

		if _, ok := (*templateMap)["#intent#"]; ok {
			(*templateMap)["#intent#"] = answer.Intention // TODO(TO/GA): decode it
		}
	}

	// InterviewID
	{
		interview, itvErr := model.QueryInterviewByID(interviewID)

		if _, ok := (*templateMap)["#interview#"]; ok {
			if itvErr != nil {
				return itvErr
			}
			(*templateMap)["#interview#"] = interview.Description // TODO(TO/GA): what the fuck?
		}

		if _, ok := (*templateMap)["#time#"]; ok {
			if itvErr != nil {
				return itvErr
			}
			(*templateMap)["#time"] = interview.StartTime.String() // TODO(TO/GA): start time or end time?
		}

		if _, ok := (*templateMap)["#location#"]; ok {
			if itvErr != nil {
				return itvErr
			}
			(*templateMap)["#location"] = interview.Location
		}
	}

	// DepartmentID
	{
		department, departErr := model.QueryDepartmentById(departmentID)
		if _, ok := (*templateMap)["#depart#"]; ok {
			if departErr != nil {
				return departErr
			}
			(*templateMap)["#depart#"] = department.Name // TODO(TO/GA): cross interview
		}

		if _, ok := (*templateMap)["#association#"]; ok {
			if departErr != nil {
				return departErr
			}
			// TODO(TO/GA): preload
			organization, orgErr := model.QueryOrganizationById(department.OrganizationID)
			if orgErr != nil {
				return orgErr
			}
			(*templateMap)["#association#"] = organization.Name // TODO(TO/GA): test
		}
	}

	// WTF
	{
		if _, ok := (*templateMap)["#url#"]; ok {
			// TODO(TO/GA): finish
		}
	}

	return nil
}

func SendMessage(messageRequest model.MessageRequest) error {
	// fetch info
	_, msgTplErr := model.QueryMessageTemplateById(messageRequest.MessageTemplateID)
	if msgTplErr != nil {
		if errors.Is(msgTplErr, gorm.ErrRecordNotFound) {
			return errors.New("message template not found")
		}
		return model.ErrInternalError
	}
	answer, ansErr := model.QueryAnswerById(messageRequest.AnswerID)
	if ansErr != nil {
		if errors.Is(ansErr, gorm.ErrRecordNotFound) {
			return errors.New("answer not found")
		}
		return model.ErrInternalError
	}
	receiver, recvErr := model.QueryUserByZJUid(answer.ZJUid)
	if recvErr != nil {
		return model.ErrInternalError
	}

	// TODO(TO/GA): call sms serview for template
	var templateMap map[string]string

	// generate map
	mapErr := GenerateMap(&templateMap, answer, messageRequest.DepartmentID, messageRequest.InterviewID)
	if mapErr != nil {
		if errors.Is(ansErr, gorm.ErrRecordNotFound) {
			return errors.New("fill placeholders fail due to the lack of information")
		}
		return model.ErrInternalError
	}

	// TODO(TO/GA): call sms service to send
	var cost float32
	var IDInSMSService uint

	// update db
	message := &model.Message{
		IDInSMSService: IDInSMSService,
		DepartmentID:   messageRequest.DepartmentID,
		ReceiverID:     receiver.ID,
		Cost:           cost,
	}
	createErr := model.CreateMessage(message)
	if createErr != nil {
		logrus.Errorf("message sent but failed to insert into database, %+v", message)
		return nil
	}

	// update cost
	department, departErr := model.QueryDepartmentById(messageRequest.DepartmentID)
	if departErr != nil {
		logrus.Errorf("fail to update department's message(ID=%v) cost, %v\n", message.ID, departErr.Error())
	}
	departErr = model.UpdateDepartmentById(&model.Department{ID: department.ID, MessageCost: department.MessageCost + cost})
	if departErr != nil {
		logrus.Errorf("fail to update department's message(ID=%v) cost, %v\n", message.ID, departErr.Error())
	}
	// TODO(TO/GA): preload
	organization, orgErr := model.QueryOrganizationById(department.OrganizationID)
	if orgErr != nil {
		logrus.Errorf("fail to update organization's message(ID=%v) cost, %v\n", message.ID, orgErr.Error())
	}
	organization.MessageCost += cost
	orgErr = model.UpdateOrganizationById(&model.Organization{ID: organization.ID, MessageCost: organization.MessageCost + cost})
	if orgErr != nil {
		logrus.Errorf("fail to update organization's message(ID=%v) cost, %v\n", message.ID, orgErr.Error())
	}

	return nil
}

func GetAllMessageSign() ([]*model.AllMessageSignAPI, error) {
	// TODO(TO/GA): call sms service
	return nil, nil
}

func AddMessageTemplate(oid uint, messageTemplateRequest *model.MessageTemplateRequest) error {
	// TODO(TO/GA): call sms service

	var IDInSMSService uint
	messageTemplate := &model.MessageTemplate{
		Title:          messageTemplateRequest.Title,
		IDInSMSService: IDInSMSService,
		OrganizationID: oid,
	}
	msgTplErr := model.CreateMessageTemplate(messageTemplate)
	if msgTplErr != nil {
		return model.ErrInternalError
	}
	return nil
}

func UpdateMessageTemplate(tid uint, messageTemplateRequest *model.MessageTemplateRequest) error {
	// fetch
	messageTemplate, msgTplErr := model.QueryMessageTemplateById(tid)
	if msgTplErr != nil {
		if errors.Is(msgTplErr, gorm.ErrRecordNotFound) {
			return errors.New("message template not found")
		}
		return model.ErrInternalError
	}

	// TODO(TO/GA): call sms service

	// update db
	messageTemplate.Title = messageTemplateRequest.Title
	msgTplErr = model.UpdateMessageTemplateById(messageTemplate)
	if msgTplErr != nil {
		logrus.Errorf("fail to update message template(ID=%v, Title=%v): %v\n", messageTemplate.ID, messageTemplateRequest.Title, msgTplErr.Error())
	}
	return nil
}

func GetMessageTemplate(tidInSMSService uint) ([]model.MessageTemplateAPI, error) {
	// TODO(TO/GA): finish it
	return nil, nil
}

func GetAllMessageTemplate(oid uint) ([]model.AllMessageTemplateAPI, error) {
	// TODO(TO/GA): finish it
	return nil, nil
}
