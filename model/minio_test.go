package model

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermanentObjectCreateAndDownload(t *testing.T) {
	ctx := context.Background()
	postUrl, policy, uuid, err := CreateObject(ctx, "foobar")
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
	assert.Equal(t, resp.StatusCode, 204)
	if err != nil {
		errorInfo, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(errorInfo)
		t.Errorf("Fail to create object")
	}

	if err = SealObject(ctx, uuid); err != nil {
		t.Errorf("Fail to get object")
	}
	getUrl, err := GetObject(ctx, uuid)
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
	assert.Equal(t, gotData, data)
}
