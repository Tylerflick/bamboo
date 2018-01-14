package bamboo

import "net/http"

type buildService interface {
	GetBuildResults(max int) Results
	GetProjects() Projects
	GetPlans() Plans
}

// RequestHandler defines an interface create and processing http requests
type RequestHandler interface {
	CreateRequest(requestType string, url string) *http.Request
	ProcessRequest(req *http.Request) string
}
