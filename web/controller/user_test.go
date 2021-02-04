package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserApi(t *testing.T) {
	t.Parallel()
	for _, v := range testCases {
		v := v  // for fear of the errors caused by go-routines
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			req := test.CreateRequest("GET", v.req.urlPath + v.req.urlQuery, nil)
			resp := test.CreateResponse(req, e)
			assert.Equal(t, v.resp.statusCode, resp.StatusCode)
			// TODO: check whether the struct (unmarshalled from JSON string in HTTP response) is expected
		})
	}
}

type req struct {
	urlPath    string
	urlQuery   string
}
type resp struct {
	statusCode int
	jsonStruct interface{}  // TODO: maybe we need to change the type of this
}

/* TODO: add more test cases */
var testCases = []struct {
	name string
	req  req
	resp resp
} {
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
			urlQuery: "?oid=100&did=1",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneNonExistingDepartmentFromOneExistingOrganization",
		req: req{
			urlPath: "/api/organization/department",
			urlQuery: "?oid=1&did=100",
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
	},
}
