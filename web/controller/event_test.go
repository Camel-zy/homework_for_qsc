package controller

import (
	"net/http"
	"testing"

	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/stretchr/testify/assert"
)

func TestEventApi(t *testing.T) {
	t.Parallel()
	for _, v := range eventTestCases {
		v := v // for fear of the errors caused by go-routines
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			Req := utils.CreateRequest("GET", v.req.urlPath+v.req.urlQuery, nil)
			Resp := utils.CreateResponse(Req, e)
			assert.Equal(t, v.resp.statusCode, Resp.StatusCode)
			// TODO: check whether the struct (unmarshalled from JSON string in HTTP Response) is expected
		})
	}
}

var eventTestCases = []testCase{
	{
		name: "GetOneExistingEvent",
		req: req{
			urlPath:  "/api/event",
			urlQuery: "?eid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneNonExistingEvent",
		req: req{
			urlPath:  "/api/event",
			urlQuery: "?eid=100",
		},
		resp: resp{
			statusCode: http.StatusForbidden,
		},
	}, {
		name: "GetOneExistingEventFromOneExistingOrganization",
		req: req{
			urlPath:  "/api/organization/event",
			urlQuery: "?oid=1&eid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingEventFromOneNonExistingOrganization",
		req: req{
			urlPath:  "/api/organization/event",
			urlQuery: "?oid=100&eid=1",
		},
		resp: resp{
			statusCode: http.StatusForbidden,
		},
	}, {
		name: "GetOneNonExistingEventFromOneExistingOrganization",
		req: req{
			urlPath:  "/api/organization/event",
			urlQuery: "?oid=1&eid=100",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequest",
		req: req{
			urlPath:  "/api/organization/event",
			urlQuery: "?oid=1",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	}, {
		name: "GetAllEventsFromOneExistingOrganization",
		req: req{
			urlPath:  "/api/organization/event/all",
			urlQuery: "?oid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	},
	{
		name: "GetAllEventsFromOneNoneExistingOrganization",
		req: req{
			urlPath:  "/api/organization/event/all",
			urlQuery: "?oid=100",
		},
		resp: resp{
			statusCode: http.StatusForbidden,
		},
	}, {
		name: "BadRequest",
		req: req{
			urlPath:  "/api/organization/event/all",
			urlQuery: "?oid=AStupidStringThatMayCrashTheService",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	},
}
