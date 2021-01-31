/*
This file seems to have no relation with any test procedure,
but the functions in this file will be called only during the unit test.
 */
package test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

/*
Mock a HTTP request sent by front-end
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
