package database

import (
	"git.zjuqsc.com/rop/rop-back-neo/database/model"
)

/*
This function will be called by the main test procedure
And should not be called out of test
*/
func CreateRowsForTest() {
	migration := []interface{} {
		&model.Organization {
			ID: uint(1),
			Name: "QSC",
			Description: "A super cool organization",
		},
		&model.Organization {
			Name: "XueGongBu",
			Description: "2333333",
		},
		&model.Department {
			Name: "Tech",
			OrganizationID: uint(1),
			Description: "Geeks! Have fun!",
		},
		&model.Department {
			Name: "Design",
			OrganizationID: uint(1),
			Description: "Wow that's beautiful!",
		},
	}

	for _, v := range migration {
		result := DB.Create(v)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
