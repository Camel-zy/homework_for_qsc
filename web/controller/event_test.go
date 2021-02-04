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
			Req := test.CreateRequest("GET", v.Req.urlPath + v.Req.urlQuery, nil)
			Resp := test.CreateResponse(Req, e)
			assert.Equal(t, v.Resp.statusCode, Resp.StatusCode)
			// TODO: check whether the struct (unmarshalled from JSON string in HTTP Response) is expected
		})
	}
}

type Req struct {
	urlPath    string
	urlQuery   string
}
type Resp struct {
	statusCode int
	jsonStruct interface{}  // TODO: maybe we need to change the type of this
}

var testCases = []struct {
	name string
	Req  Req
	Resp Resp
} {
	{
		name: "GetOneExistingEventFromOneExistingOrganization",
		Req: Req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=1&eid=1",
		},
		Resp: Resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingEventFromOneNonExistingOrganization",
		Req: Req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=100&eid=1",
		},
		Resp: Resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneNonExistingEventFromOneExistingOrganization",
		Req: Req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=1&eid=100",
		},
		Resp: Resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequestForOrganization",
		Req: Req{
			urlPath: "/api/organization/event",
			urlQuery: "?oid=1",
		},
		Resp: Resp{
			statusCode: http.StatusBadRequest,
		},
	},
	{
		name: "GetOneExistingInterviewFromOneExistingEvent",
		Req: Req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=1&iid=1",
		},
		Resp: Resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingInterviewFromOneNonExistingEvent",
		Req: Req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=100&iid=1",
		},
		Resp: Resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneNonExistingInterviewFromOneExistingEvent",
		Req: Req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=1&iid=100",
		},
		Resp: Resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequestForEvent",
		Req: Req{
			urlPath: "/api/event/interview",
			urlQuery: "?eid=1",
		},
		Resp: Resp{
			statusCode: http.StatusBadRequest,
		},
	},
}
