package bamboo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Service is a struct which exposes functinality for talking to the Attlasian Bamboo REST API
type Service struct {
	Username       string
	Password       string
	BaseURL        string
	requestHandler RequestHandler
}

// New initializes a new service struct
func New(userName string, password string, baseURL string, requestHandler RequestHandler) Service {
	b := Service{
		Username: userName,
		Password: password,
		BaseURL:  baseURL,
	}
	if requestHandler == nil {
		b.requestHandler = b
	} else {
		b.requestHandler = requestHandler
	}
	return b
}

// GetBuildResults retrieves the last n(max) build results
func (b Service) GetBuildResults(max int) Results {
	endpoint := b.BaseURL + ApiUrl + LatestResults
	req := b.requestHandler.CreateRequest("GET", endpoint)
	if max > 0 {
		values := req.URL.Query()
		values.Add(MaxResultsKey, strconv.Itoa(max))
		req.URL.RawQuery = values.Encode()
	}
	req.SetBasicAuth(b.Username, b.Password)
	body := b.requestHandler.ProcessRequest(req)
	var jsonData Results
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

// GetPlanResult retrieves the build result for a particular plan
func (b Service) GetPlanResult(planKey string) PlanResult {
	endpoint := b.BaseURL + ApiUrl + "/latest/result/" + planKey
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	values := req.URL.Query()
	req.URL.RawQuery = values.Encode()
	body := b.requestHandler.ProcessRequest(req)
	var jsonData PlanResult
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

// GetProjects retrieves all projects
func (b Service) GetProjects() Projects {
	endpoint := b.BaseURL + ApiUrl + ProjectEndpoint
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	body := b.requestHandler.ProcessRequest(req)
	var jsonData Projects
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

// GetPlans retrieves the last n(max) plans
func (b Service) GetPlans(max int) Plans {
	endpoint := b.BaseURL + ApiUrl + PlanEndpoint
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	if max > 0 {
		values := req.URL.Query()
		values.Add(MaxResultsKey, strconv.Itoa(max))
		req.URL.RawQuery = values.Encode()
	}
	body := b.requestHandler.ProcessRequest(req)
	var jsonData Plans
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

// GetPlanDetails retrieves the PlanDetails for an individual plan
func (b Service) GetPlanDetails(planKey string) PlanDetails {
	endpoint := b.BaseURL + ApiUrl + PlanEndpoint + "/" + planKey
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	body := b.requestHandler.ProcessRequest(req)
	var jsonData PlanDetails
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

// GetPlanBranches retrieves all child branches of a particular plan
func (b Service) GetPlanBranches(planKey string) PlanBranches {
	endpoint := b.BaseURL + ApiUrl + fmt.Sprintf(ListPlanBranchesEndpoint, planKey)
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	body := b.requestHandler.ProcessRequest(req)
	var jsonData PlanBranches
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

// CreateRequest builds a json based http request
func (b Service) CreateRequest(requestType string, url string) *http.Request {
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

// ProcessRequest excecutes a http request and returns the body as a string
func (b Service) ProcessRequest(req *http.Request) string {
	client := http.Client{}
	results, _ := client.Do(req)
	return b.getBody(results)
}

func (b Service) getBody(response *http.Response) string {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
