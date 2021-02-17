package utils

import (
	"strings"

	"git.zjuqsc.com/rop/rop-back-neo/model"
)

func GenerateMap(messageRequest model.MessageRequest) (map[string]string, error) {
	messageTemplate, msgTplErr := model.QueryMessageTemplateById(messageRequest.MessageTemplateID)
	if msgTplErr != nil {
		return nil, msgTplErr
	}

	ret := make(map[string]string)

	// ReceiverID
	{
		receiver, recvErr := model.QueryUserById(messageRequest.ReceiverID)
		if msgTplErr != nil {
			return nil, recvErr
		}

		if strings.Contains(messageTemplate.Text, "#name#") {
			ret["#name#"] = receiver.Name
		}

		if strings.Contains(messageTemplate.Text, "#stuid#") {
			ret["#stuid#"] = receiver.ZJUid
		}

		// TODO(TO/GA): move to FormID section
		if strings.Contains(messageTemplate.Text, "#phone#") {
			ret["#phone#"] = receiver.Mobile
		}
	}

	// InterviewID
	{
		interview, itvErr := model.QueryInterviewByID(messageRequest.InterviewID)

		if strings.Contains(messageTemplate.Text, "#interview#") {
			if itvErr != nil {
				return nil, itvErr
			}
			ret["#interview#"] = interview.Description // TODO(TO/GA): what the fuck?
		}

		if strings.Contains(messageTemplate.Text, "#time#") {
			if itvErr != nil {
				return nil, itvErr
			}
			ret["#time"] = interview.StartTime.String() // TODO(TO/GA): start time or end time?
		}

		if strings.Contains(messageTemplate.Text, "#location#") {
			if itvErr != nil {
				return nil, itvErr
			}
			ret["#location"] = interview.Location
		}

		// DepartmentID
		{
			department, departErr := model.QueryDepartmentById(interview.DepartmentID)
			if strings.Contains(messageTemplate.Text, "#depart#") {
				if itvErr != nil {
					return nil, itvErr
				}
				if departErr != nil {
					return nil, departErr
				}
				ret["#depart#"] = department.Name
			}

			if strings.Contains(messageTemplate.Text, "#association#") {
				if itvErr != nil {
					return nil, itvErr
				}
				if departErr != nil {
					return nil, departErr
				}
				organization, orgErr := model.QueryOrganizationById(department.OrganizationID)
				// TODO(TO/GA): looks like department.Organization doesn't update automatically
				if orgErr != nil {
					return nil, orgErr
				}
				ret["#association#"] = organization.Name // TODO(TO/GA): test
			}
		}
	}

	// FormID
	{
		if strings.Contains(messageTemplate.Text, "#intent#") {

		}

		if strings.Contains(messageTemplate.Text, "#url#") {

		}
	}

	return ret, nil
}
