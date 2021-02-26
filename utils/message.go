package utils

import (
	"errors"
	"fmt"
	"strings"

	"git.zjuqsc.com/rop/rop-back-neo/model"
	"git.zjuqsc.com/rop/rop-back-neo/rpc"
	SMSService "git.zjuqsc.com/rop/rop-sms/gRPC"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetMessageBalance() (float32, error) {
	reply, replyErr := rpc.Sms.UserBalance(&SMSService.UsrReq{})
	if replyErr != nil {
		return 0, replyErr
	}
	return reply.Balance, nil
}

var roundName = []string{"0", "公海", "一面", "二面", "三面", "四面", "五面", "六面", "七面", "八面", "九面"}

func GenerateText(messageTemplateID uint, interviewee *model.Interviewee, answer *model.Answer) (*string, error) {
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
		department, departErr := model.QueryDepartmentById(interviewee.DepartmentID)
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
		// interview, itvErr := model.QueryInterviewByID(interviewee.)
		interview, itvErr := model.QueryInterviewByIntervieweeAndRound(interviewee.ID, interviewee.Round)

		if strings.Contains(templateText, "#interview#") {
			if itvErr != nil {
				return nil, itvErr
			}
			templateText = strings.ReplaceAll(templateText, "#interview#", roundName[interview.Round])
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

func SendMessage(vid uint, messageTemplateID uint) (*string, error) {
	// fetch info
	interviewee, intervieweeErr := model.QueryIntervieweeById(vid)
	if intervieweeErr != nil {
		if errors.Is(intervieweeErr, gorm.ErrRecordNotFound) {
			return nil, errors.New("interviewee not found")
		}
		return nil, model.ErrInternalError
	}
	answer, ansErr := model.QueryAnswerByID(interviewee.AnswerID)
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
	text, textErr := GenerateText(messageTemplateID, interviewee, answer)
	if textErr != nil {
		if errors.Is(ansErr, gorm.ErrRecordNotFound) {
			return nil, errors.New("fill placeholders fail due to the lack of information")
		}
		logrus.Errorf("GenerateText failed(vid=%v, mid=%v):%v", vid, messageTemplateID, textErr)
		return nil, model.ErrInternalError
	}

	fmt.Printf("send to %v\n%v\n", answer.Mobile, *text) // TODO(TO/GA): delete it

	reply, replyErr := rpc.Sms.SendMsgByText(&SMSService.MsgReq{
		Mobile: answer.Mobile,
		Text:   *text,
	})
	if replyErr != nil {
		return nil, replyErr // TODO(TO/GA): error handling
	}

	// update db
	message := &model.Message{
		IDInSMSService: uint(reply.ID),
		DepartmentID:   interviewee.DepartmentID,
		ReceiverID:     receiver.ID,
		Cost:           reply.Fee,
	}
	createErr := model.CreateMessage(message)
	if createErr != nil {
		logrus.Errorf("message sent but failed to insert into database, %+v\n", message)
		return text, nil
	}

	// update cost
	department, departErr := model.QueryDepartmentById(interviewee.DepartmentID)
	if departErr != nil {
		logrus.Errorf("fail to update department's message(ID=%v) cost\n", message.ID)
	}
	departErr = model.UpdateDepartmentById(&model.Department{ID: department.ID, MessageCost: department.MessageCost + reply.Fee})
	if departErr != nil {
		logrus.Errorf("fail to update department's message(ID=%v) cost\n", message.ID)
	}
	// TODO(TO/GA): preload
	organization, orgErr := model.QueryOrganizationById(department.OrganizationID)
	if orgErr != nil {
		logrus.Errorf("fail to update organization's message(ID=%v) cost\n", message.ID)
	}
	organization.MessageCost += reply.Fee
	orgErr = model.UpdateOrganizationById(&model.Organization{ID: organization.ID, MessageCost: organization.MessageCost + reply.Fee})
	if orgErr != nil {
		logrus.Errorf("fail to update organization's message(ID=%v) cost\n", message.ID)
	}

	return text, nil
}
