package tests

import (
	"net/http"
)

type MockRequestHandler struct {
	BaseUrl      string
	RequestUrl   string
	ResponseBody string
}

func (m *MockRequestHandler) CreateRequest(requestType string, url string) *http.Request {
	m.RequestUrl = url
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

func (m *MockRequestHandler) ProcessRequest(req *http.Request) string {
	return m.ResponseBody
}
