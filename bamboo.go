package bamboo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type BambooService struct {
	Username       string
	Password       string
	BaseUrl        string
	requestHandler RequestHandler
}

func NewBambooService(userName string, password string, baseUrl string, requestHandler RequestHandler) BambooService {
	b := BambooService{
		Username: userName,
		Password: password,
		BaseUrl:  baseUrl,
	}
	if requestHandler == nil {
		b.requestHandler = b
	} else {
		b.requestHandler = requestHandler
	}
	return b
}

func (b BambooService) GetBuildResults(max int) Results {
	endpoint := b.BaseUrl + ApiUrl + LatestResults
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
func (b BambooService) GetPlanResult(id string) PlanResult {
	endpoint := b.BaseUrl + ApiUrl + "/latest/result/" + id
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	values := req.URL.Query()
	req.URL.RawQuery = values.Encode()
	body := b.requestHandler.ProcessRequest(req)
	var jsonData PlanResult
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b BambooService) GetProjects() Projects {
	endpoint := b.BaseUrl + ApiUrl + ProjectEndpoint
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	body := b.requestHandler.ProcessRequest(req)
	var jsonData Projects
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b BambooService) GetPlans(max int) Plans {
	endpoint := b.BaseUrl + ApiUrl + PlanEndpoint
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

func (b BambooService) GetPlanDetails(id string) PlanDetails {
	endpoint := b.BaseUrl + ApiUrl + PlanEndpoint + "/" + id
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	body := b.requestHandler.ProcessRequest(req)
	var jsonData PlanDetails
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b BambooService) GetPlanBranches(buildKey string) PlanBranches {
	endpoint := b.BaseUrl + ApiUrl + fmt.Sprintf(ListPlanBranchesEndpoint, buildKey)
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)
	body := b.requestHandler.ProcessRequest(req)
	var jsonData PlanBranches
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b BambooService) CreateRequest(requestType string, url string) *http.Request {
	req, _ := http.NewRequest(requestType, url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

func (b BambooService) ProcessRequest(req *http.Request) string {
	client := http.Client{}
	results, _ := client.Do(req)
	return b.getBody(results)
}

func (b BambooService) getBody(response *http.Response) string {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
