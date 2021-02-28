package model

/*
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
		_ = w.WriteField(k, v)
	}
	str := "gg is 渣男"
	data := []byte(str)
	fw, _ := w.CreateFormFile("file", "test.txt")
	_, _ = fw.Write(data)
	_ = w.Close()
	req, err := http.NewRequest("POST", postUrl.String(), &b)
	assert.NotNil(t, err)
	req.Header.Add("Content-Type", w.FormDataContentType())

	resp, err := hc.Do(req)
	assert.Equal(t, resp.StatusCode, http.StatusNoContent)
	if err != nil {
		errorInfo, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(errorInfo)
		t.Errorf("Fail to create object")
	}

	assert.NotNil(t, SealObject(ctx, uuid))
	getUrl, err := GetObject(ctx, uuid)
	assert.NotNil(t, err)
	req, err = http.NewRequest("GET", getUrl.String(), nil)
	assert.NotNil(t, err)
	resp, err = hc.Do(req)
	assert.NotNil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	gotData, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, gotData, data)
}

 */
