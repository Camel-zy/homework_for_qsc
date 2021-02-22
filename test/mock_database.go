package test

import (
	"git.zjuqsc.com/rop/rop-back-neo/model"
)

/*
This function will be called by the main test procedure
And should not be called out of test
*/
func CreateDatabaseRows() {
	migration := []interface{} {
		&model.User{
			ID: uint(1),
			Name: "testing_user",
			ZJUid: "3190102333",
		},
		&model.OrganizationHasUser{
			ID: uint(1),
			UserId: uint(1),
			OrganizationId: uint(1),
			DepartmentId: uint(1),
		},
		&model.Organization{
			ID: uint(1),
			Name: "QSC",
			Description: "A super cool organization",
		},
		&model.Organization{
			ID: uint(2),
			Name: "XueGongBu",
			Description: "2333333",
		},
		&model.Department{
			ID: uint(1),
			Name: "Tech",
			OrganizationID: uint(1),
			Description: "Geeks! Have fun!",
		},
		&model.Department{
			ID: uint(2),
			Name: "Design",
			OrganizationID: uint(1),
			Description: "Wow that's beautiful!",
		},
		&model.Department{
			ID: uint(3),
			Name: "StudentManagement",
			OrganizationID: uint(2),
			Description: "emm",
		},
		&model.Event{
			ID: uint(1),
			Name: "Fall",
			OrganizationID: uint(1),
		},
		&model.Event{
			ID: uint(2),
			Name: "Spring",
			OrganizationID: uint(1),
		},
		&model.Interview{
			ID: uint(1),
			Name: "Round_One",
			EventID: uint(1),
		},
		&model.Interview{
			ID: uint(2),
			Name: "Round_Two",
			EventID: uint(1),
		},
	}

	for _, v := range migration {
		result := model.CreateRow(v)
		if result != nil {
			panic(result.Error)
		}
	}
}
