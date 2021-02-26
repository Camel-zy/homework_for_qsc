package test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"

	"git.zjuqsc.com/rop/rop-back-neo/conf"
	"git.zjuqsc.com/rop/rop-back-neo/model"
	"gorm.io/driver/postgres"
)

func testInit() {
	conf.Init()
	model.Connect(postgres.Open(conf.GetDatabaseLoginInfo()))
	model.CreateTables()
	model.ConnectObjectStorage()
}
func TestPermanentObjectCreateAndDownload(t *testing.T) {
	testInit()
	ctx := context.Background()
	postUrl, policy, uuid, err := model.CreateObject(ctx, "foobar")
	if err != nil {
		t.Errorf("Fail to create object")
	}
	hc := http.Client{}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for k, v := range policy {
		w.WriteField(k, v)
	}
	str := "gg is 渣男"
	data := []byte(str)
	fw, _ := w.CreateFormFile("file", "test.txt")
	fw.Write(data)
	w.Close()
	req, err := http.NewRequest("POST", postUrl.String(), &b)
	if err != nil {
		t.Errorf("Fail to create object")
	}
	req.Header.Add("Content-Type", w.FormDataContentType())

	resp, err := hc.Do(req)
	if err != nil || resp.StatusCode != 204 {
		var errorInfo []byte
		_, _ = resp.Body.Read(errorInfo)
		fmt.Println(errorInfo)
		t.Errorf("Fail to create object")
	}
	err = model.SealObject(ctx, uuid)
	if err != nil {
		t.Errorf("Fail to get object")
	}
	getUrl, err := model.GetObject(ctx, uuid)
	if err != nil {
		t.Errorf("Fail to get object")
	}
	req, err = http.NewRequest("GET", getUrl.String(), nil)
	if err != nil {
		t.Errorf("Fail to get object")
	}
	resp, err = hc.Do(req)
	if err != nil || resp.StatusCode != 200 {
		t.Errorf("Fail to get object")
	}
	gotData, _ := ioutil.ReadAll(resp.Body)
	if !bytes.Equal(gotData, data) {
		t.Errorf("Fail to get object")
	}
}
