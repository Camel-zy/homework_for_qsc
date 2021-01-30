package web

import (
	"git.zjuqsc.com/rop/rop-back-neo/database"
	"git.zjuqsc.com/rop/rop-back-neo/web/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	/* open a sqlite in-memory database */
	database.Connect(sqlite.Open("file::memory:?cache=shared"))
	database.CreateTables()

	database.CreateRowsForTest()
	InitWebFramework(true)

	os.Exit(m.Run())
}

func TestController(t *testing.T) {
	for _, v := range testCases {
		v := v
		req := utils.CreateRequest("GET", v.req.urlPath + v.req.urlQuery, nil)
		resp := utils.CreateResponse(req, e)
		assert.Equal(t, v.resp.statusCode, resp.StatusCode)
		// TODO: check whether the struct (unmarshalled from JSON string in HTTP response) is expected
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
		name: "Get One Department from Organization",
		req: req{
			urlPath: "/api/organization/department",
			urlQuery: "?oid=1&did=1",
		},
		resp: resp{
			statusCode: http.StatusOK,
		},
	},
}