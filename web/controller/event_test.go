package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestEventApi(t *testing.T) {
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

var testCases = []struct {
	name string
	req  req
	resp resp
} {
	{
		name: "GetOneExistingEventFromOneExistingOrganization",
		req: req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=1&eid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingEventFromOneNonExistingOrganization",
		req: req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=100&eid=1",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneNonExistingEventFromOneExistingOrganization",
		req: req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=1&eid=100",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequestForOrganization",
		req: req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=1",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	},
	{
		name: "GetOneExistingInterviewFromOneExistingEvent",
		req: req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=1&iid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingInterviewFromOneNonExistingEvent",
		req: req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=100&iid=1",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneNonExistingInterviewFromOneExistingEvent",
		req: req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=1&iid=100",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequestForEvent",
		req: req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=1",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	},
}
