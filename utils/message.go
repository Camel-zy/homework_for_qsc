package utils

import (
	"errors"
	"strings"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetMessageBalance() (float32, error) {
	//TODO(TO/GA): call sms service
	var cost float32

	return cost, nil
}

func GenerateText(messageTemplateID uint, answer *model.Answer, departmentID uint, interviewID uint) (*string, error) {
	templateText := model.MessageTemplate[messageTemplateID]

	// answer
	{
		if strings.Contains(templateText, "#name#") {
			templateText = strings.ReplaceAll(templateText, "#name#", answer.Name)
		}

		if strings.Contains(templateText, "#stuid#") {
			templateText = strings.ReplaceAll(templateText, "#stuid#", answer.ZJUid)
		}

		if strings.Contains(templateText, "#intent#") {
			// templateText = strings.ReplaceAll(templateText, "#intent#", answer.Intention) // TODO(TO/GA): decode it
		}
	}

	// EventID
	{
		event, eventErr := model.QueryEventByID(answer.EventID)
		if eventErr != nil {
			return nil, eventErr
		}
		if strings.Contains(templateText, "#event#") {
			templateText = strings.ReplaceAll(templateText, "#event#", event.Name)
		}
	}

	// DepartmentID
	{
		department, departErr := model.QueryDepartmentById(departmentID)
		if strings.Contains(templateText, "#depart#") {
			if departErr != nil {
				return nil, departErr
			}
			templateText = strings.ReplaceAll(templateText, "#depart#", department.Name) // TODO(TO/GA): cross interview
		}

		if strings.Contains(templateText, "#association#") {
			if departErr != nil {
				return nil, departErr
			}
			// TODO(TO/GA): preload
			organization, orgErr := model.QueryOrganizationById(department.OrganizationID)
			if orgErr != nil {
				return nil, orgErr
			}
			templateText = strings.ReplaceAll(templateText, "#association#", organization.Name)
		}
	}

	// InterviewID
	if messageTemplateID == 1 || messageTemplateID == 2 {
		interview, itvErr := model.QueryInterviewByID(interviewID)

		if strings.Contains(templateText, "#interview#") {
			if itvErr != nil {
				return nil, itvErr
			}
			templateText = strings.ReplaceAll(templateText, "#interview#", interview.Name)
		}

		if strings.Contains(templateText, "#time#") {
			if itvErr != nil {
				return nil, itvErr
			}
			templateText = strings.ReplaceAll(templateText, "#time#", interview.StartTime.String())
		}

		if strings.Contains(templateText, "#location#") {
			if itvErr != nil {
				return nil, itvErr
			}
			templateText = strings.ReplaceAll(templateText, "#location#", interview.Location)
		}
	}

	// WTF
	{
		if strings.Contains(templateText, "#url#") {
			// TODO(TO/GA): finish
		}
	}

	return &templateText, nil
}

func SendMessage(messageRequest *model.MessageRequest, messageTemplateID uint) (*string, error) {
	// fetch info
	answer, ansErr := model.QueryAnswerById(messageRequest.AnswerID)
	if ansErr != nil {
		if errors.Is(ansErr, gorm.ErrRecordNotFound) {
			return nil, errors.New("answer not found")
		}
		return nil, model.ErrInternalError
	}
	receiver, recvErr := model.QueryUserByZJUid(answer.ZJUid)
	if recvErr != nil {
		return nil, model.ErrInternalError
	}

	// generate map
	text, textErr := GenerateText(messageTemplateID, answer, messageRequest.DepartmentID, messageRequest.InterviewID)
	if textErr != nil {
		if errors.Is(ansErr, gorm.ErrRecordNotFound) {
			return nil, errors.New("fill placeholders fail due to the lack of information")
		}
		return nil, model.ErrInternalError
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
		logrus.Errorf("message sent but failed to insert into database, %+v\n", message)
		return text, nil
	}

	// update cost
	department, departErr := model.QueryDepartmentById(messageRequest.DepartmentID)
	if departErr != nil {
		logrus.Errorf("fail to update department's message(ID=%v) cost\n", message.ID)
	}
	departErr = model.UpdateDepartmentById(&model.Department{ID: department.ID, MessageCost: department.MessageCost + cost})
	if departErr != nil {
		logrus.Errorf("fail to update department's message(ID=%v) cost\n", message.ID)
	}
	// TODO(TO/GA): preload
	organization, orgErr := model.QueryOrganizationById(department.OrganizationID)
	if orgErr != nil {
		logrus.Errorf("fail to update organization's message(ID=%v) cost\n", message.ID)
	}
	organization.MessageCost += cost
	orgErr = model.UpdateOrganizationById(&model.Organization{ID: organization.ID, MessageCost: organization.MessageCost + cost})
	if orgErr != nil {
		logrus.Errorf("fail to update organization's message(ID=%v) cost\n", message.ID)
	}

	return text, nil
}
