package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDepartmentApi(t *testing.T) {
	t.Parallel()
	for _, v := range testCases {
		v := v  // for fear of the errors caused by go-routines
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			req := utils.CreateRequest("GET", v.req.urlPath + v.req.urlQuery, nil)
			resp := utils.CreateResponse(req, e)
			assert.Equal(t, v.resp.statusCode, resp.StatusCode)
			// TODO: check whether the struct (unmarshalled from JSON string in HTTP response) is expected
		})
	}
}

var testCases = []testCase{
	{
		name: "GetOneExistingDepartmentFromOneExistingOrganization",
		req: req{
			urlPath: "/api/organization/department",
			urlQuery: "?oid=1&did=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingDepartmentFromOneNonExistingOrganization",
		req: req{
			urlPath: "/api/organization/department",
			urlQuery: "?oid=2147483647&did=1",
		},
		resp: resp{
			statusCode: http.StatusForbidden,
		},
	}, {
		name: "GetOneNonExistingDepartmentFromOneExistingOrganization",
		req: req{
			urlPath: "/api/organization/department",
			urlQuery: "?oid=1&did=2147483647",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneExistingDepartmentFromOneExistingOrganizationButNotIn",
		req: req{
			urlPath: "/api/organization/department",
			urlQuery: "?oid=1&did=3",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequest",
		req: req{
			urlPath: "/api/organization/department",
			urlQuery: "?oid=1",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	}, {
		name: "GetAllDepartmentsFromOneExistingOrganization",
		req: req{
			urlPath: "/api/organization/department/all",
			urlQuery: "?oid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	},
	{
		name: "GetAllDepartmentsFromOneNoneExistingOrganization",
		req: req{
			urlPath: "/api/organization/department/all",
			urlQuery: "?oid=2147483647",
		},
		resp: resp{
			statusCode: http.StatusForbidden,
		},
	}, {
		name: "BadRequest",
		req: req{
			urlPath: "/api/organization/department/all",
			urlQuery: "?oid=AStupidStringThatMayCrashTheService",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	},
}
