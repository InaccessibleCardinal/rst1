package ht

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	applicationJson = "application/json"
	contentType     = "Content-Type"
	authorization   = "Authorization"
	authToken       = "temp"
)

type HtFunc func(method string, url string, bodyMap map[string]any) (string, error)

func makeRequestBody(bodyMap map[string]any) (io.Reader, error) {
	bodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bodyBytes), nil
}

func setHeaders(req *http.Request) {
	req.Header.Set(contentType, applicationJson)
	req.Header.Set(authorization, authToken)
}

func makeRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	setHeaders(req)
	return req, err
}

func errorResponse(err error) (string, error) {
	println(err.Error())
	return "", err
}

func checkStatus() {

}

func MakeHttpRequest(method string, url string, bodyMap map[string]any) (string, error) {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{Timeout: timeout}
	body, err := makeRequestBody(bodyMap)
	if err != nil {
		return errorResponse(err)
	}
	req, err := makeRequest(method, url, body)
	if err != nil {
		return errorResponse(err)
	}
	res, err := client.Do(req)
	if err != nil {
		return errorResponse(err)
	}
	if res.StatusCode > 399 {
		return errorResponse(errors.New("bad status: " + strconv.Itoa(res.StatusCode)))
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errorResponse(err)
	}
	return string(responseBody), nil
}
