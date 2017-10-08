package bamboo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type bambooService struct {
	username string
	password string
	baseUrl  string
}

func (b bambooService) GetBuildResults(max int) Results {
	endpoint := b.baseUrl + apiUrl + latestResults
	req := b.createRequest("GET", endpoint)
	if max > 0 {
		req.URL.Query().Add(maxResultsKey, strconv.Itoa(max))
	}
	body := b.processRequeset(req)
	var jsonData Results
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b bambooService) GetProjects() Projects {
	endpoint := b.baseUrl + apiUrl + project
	req := b.createRequest("GET", endpoint)
	body := b.processRequeset(req)
	var jsonData Projects
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b bambooService) GetPlans() Plans {
	endpoint := b.baseUrl + apiUrl + plan
	req := b.createRequest("GET", endpoint)
	body := b.processRequeset(req)
	var jsonData Plans
	json.Unmarshal([]byte(body), &jsonData)
	return jsonData
}

func (b bambooService) createRequest(requestType string, url string) *http.Request {
	req, _ := http.NewRequest(requestType, url, nil)
	req.SetBasicAuth(b.username, b.password)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

func (b bambooService) processRequeset(req *http.Request) string {
	client := http.Client{}
	results, _ := client.Do(req)
	return b.getBody(results)
}

func (b bambooService) getBody(response *http.Response) string {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	// TODO: should this log the error or return it?
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
