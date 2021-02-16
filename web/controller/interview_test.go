package controller

import (
	"net/http"
	"testing"

	"git.zjuqsc.com/rop/rop-back-neo/utils"
	"github.com/stretchr/testify/assert"
)

func TestInterviewApi(t *testing.T) {
	t.Parallel()
	for _, v := range interviewTestCases {
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

var interviewTestCases = []testCase{
	{
		name: "GetOneExistingInterview",
		req: req{
			urlPath:  "/api/interview",
			urlQuery: "?iid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneNonExistingInterview",
		req: req{
			urlPath:  "/api/interview",
			urlQuery: "?iid=100",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneExistingInterviewFromOneExistingEvent",
		req: req{
			urlPath:  "/api/event/interview",
			urlQuery: "?eid=1&iid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	}, {
		name: "GetOneExistingInterviewFromOneNonExistingEvent",
		req: req{
			urlPath:  "/api/event/interview",
			urlQuery: "?eid=100&iid=1",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "GetOneNonExistingInterviewFromOneExistingEvent",
		req: req{
			urlPath:  "/api/event/interview",
			urlQuery: "?eid=1&iid=100",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequest",
		req: req{
			urlPath:  "/api/event/interview",
			urlQuery: "?eid=1",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	}, {
		name: "GetAllInterviewsFromOneExistingEvent",
		req: req{
			urlPath:  "/api/event/interview/all",
			urlQuery: "?eid=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	},
	{
		name: "GetAllInterviewsFromOneNoneExistingEvent",
		req: req{
			urlPath:  "/api/event/interview/all",
			urlQuery: "?eid=100",
		},
		resp: resp{
			statusCode: http.StatusNotFound,
		},
	}, {
		name: "BadRequest",
		req: req{
			urlPath:  "/api/event/interview/all",
			urlQuery: "?oid=AStupidStringThatMayCrashTheService",
		},
		resp: resp{
			statusCode: http.StatusBadRequest,
		},
	},
}
