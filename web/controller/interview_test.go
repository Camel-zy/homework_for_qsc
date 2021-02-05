package controller

import (
	"git.zjuqsc.com/rop/rop-back-neo/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInterviewApi(t *testing.T) {
	t.Parallel()
	for _, v := range interviewTestCases {
		v := v  // for fear of the errors caused by go-routines
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			Req := test.CreateRequest("GET", v.req.urlPath + v.req.urlQuery, nil)
			Resp := test.CreateResponse(Req, e)
			assert.Equal(t, v.resp.statusCode, Resp.StatusCode)
			// TODO: check whether the struct (unmarshalled from JSON string in HTTP Response) is expected
		})
	}
}

var interviewTestCases = []testCase{
	{
		name: "GetOneExistingInterviewFromOneExistingEvent",
		req: req{
			urlPath: "/api/organization/event/interview",
			urlQuery: "?eid=1&iid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingInterviewFromOneNonExistingEvent",
		req: req{
			urlPath: "/api/organization/event/interview",
			urlQuery: "?eid=100&iid=1",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneNonExistingInterviewFromOneExistingEvent",
		req: req{
			urlPath: "/api/organization/event/interview",
			urlQuery: "?eid=1&iid=100",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequestForEvent",
		req: req{
			urlPath: "/api/organization/event/interview",
			urlQuery: "?eid=1",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	},
}
