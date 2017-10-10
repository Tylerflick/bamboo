package bamboo

import (
	"encoding/json"
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

func NewBambooService(requestHandler RequestHandler) BambooService {
	b := BambooService{}
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
		values.Add(maxResultsKey, strconv.Itoa(max))
		req.URL.RawQuery = values.Encode()
	}
	req.SetBasicAuth(b.Username, b.Password)	
	body := b.requestHandler.ProcessRequest(req)
	var jsonData Results
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b BambooService) GetProjects() Projects {
	endpoint := b.BaseUrl + ApiUrl + project
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)	
	body := b.requestHandler.ProcessRequest(req)
	var jsonData Projects
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b BambooService) GetPlans() Plans {
	endpoint := b.BaseUrl + ApiUrl + plan
	req := b.requestHandler.CreateRequest("GET", endpoint)
	req.SetBasicAuth(b.Username, b.Password)	
	body := b.requestHandler.ProcessRequest(req)
	var jsonData Plans
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
	// TODO: should this log the error or return it?
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
