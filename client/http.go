package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

type SMHttp struct {
	baseUrl   *url.URL
	apiKey    string
	env       string
	timeOut   time.Duration
	client    *http.Client
	userAgent string
}

func NewSMHttp(smUrl, key, env string, timeOut time.Duration) *SMHttp {
	e := &SMHttp{}
	if len(key) == 0 {
		log.Fatalf("no api key")
	}
	e.apiKey = key
	e.env = env
	e.client = &http.Client{}
	e.timeOut = timeOut

	baseUrlP, err := url.Parse(smUrl)
	if err != nil {
		log.Fatalf("could not determine url with error %s", err.Error())
	}
	e.baseUrl = baseUrlP
	log.Printf("base: %s", baseUrlP)
	e.userAgent = "goSMAPI/1.0"
	return e
}

func (sm *SMHttp) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	//rel := &url.URL{Path: path}
	//u := c.BaseURL.ResolveReference(rel)

	u := fmt.Sprintf("%s/%s", sm.baseUrl, path)

	var buf io.ReadWriter
	if body != nil {
		jsonBody, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(jsonBody)
		//fmt.Println("buff: ", body)
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", sm.apiKey))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go SurveyMonkey Client")

	return req, nil
}

func (sm *SMHttp) Do(req *http.Request, v interface{}) (*http.Response, error) {

	ctx, _ := context.WithTimeout(context.TODO(), sm.timeOut)
	req = req.WithContext(ctx)

	resp, err := sm.client.Do(req)
	if err != nil {
		log.Printf("error sending request: %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	//fmt.Printf(" header %d", resp.StatusCode)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	//fmt.Printf("response body %s\n", string(bodyBytes))

	if resp.StatusCode == 200 {
		err = json.Unmarshal(bodyBytes, &v)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}
