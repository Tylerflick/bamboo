package tests

import (
	"github.com/Tylerflick/bamboo"
	"testing"
)

func TestGetPlansHitsCorrectEndpoint(t *testing.T) {
	baseUrl := "test.com"
	requestHandler := MockRequestHandler{}
	b := bamboo.NewBambooService("username", "password", baseUrl, &requestHandler)
	b.GetPlans()
	if requestHandler.RequestUrl != baseUrl + bamboo.ApiUrl + bamboo.PlanEndpoint {
		t.Error(
			"For", "requestHandler.RequestUrl",
			"expected", baseUrl + bamboo.ApiUrl + bamboo.PlanEndpoint,
			"got", requestHandler.RequestUrl,
		)
	}
}

func TestGetPlans(t *testing.T) {
	requestHandler := MockRequestHandler{
		BaseUrl: "test.com",
		ResponseBody: `{
			"expand": "plans",
			"link": {
				"rel": "self",
				"href": "http://myhost:8085/rest/api/latest/plan"
			},
			"plans": {
				"expand": "plan",
				"size": 2,
				"max-result": 2,
				"start-index": 0,
				"plan": [
					{
						"name": "My Plan 1",
						"key": "MYPLAN1",
						"link": {
							"rel": "self",
							"href": "http://myhost:8085/rest/api/latest/plan/MYPLAN1"
						}
					},
					{
						"name": "My Plan 2",
						"key": "MYPLAN2",
						"link": {
							"rel": "self",
							"href": "http://myhost:8085/rest/api/latest/plan/MYPLAN2"
						}
					},
					{
						"name": "My Plan 3",
						"key": "MYPLAN3",
						"link": {
							"rel": "self",
							"href": "http://myhost:8085/rest/api/latest/plan/MYPLAN3"
						}
					}
				]
			}
		}`,
	}
	b := bamboo.NewBambooService("username", "password", "test.com", &requestHandler)
	plans := b.GetPlans()
	if len(plans.Plans.Plans) != 3 {
		t.Error(
			"For", "len(plans.Plans.Plans)",
			"expected", 3,
			"got", len(plans.Plans.Plans),
		)
	}
}
