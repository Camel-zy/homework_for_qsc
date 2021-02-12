/*
The functions in this file will be called during the unit test,
and also might be called in the main procedure.
 */
package utils

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

/*
Create a HTTP request
 */
func CreateRequest(method string, path string, data interface{}) *http.Request {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(method, path, bytes.NewReader(jsonByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}

/*
Serve a HTTP request instantly by Echo, and return a HTTP response
 */
func CreateResponse(req *http.Request, e *echo.Echo) *http.Response {
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec.Result()
}
