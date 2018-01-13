package bamboo

const (
	ApiUrl = "/rest/api"
	// API ENDPOINTS
	LatestResults            = "/latest/result"
	ProjectEndpoint          = "/latest/projects"
	PlanEndpoint             = "/latest/plan/latest"
	ListPlanBranchesEndpoint = "/latest/plan/%s/branch" // buildKey might be top level plan (projectKey-planKey) or job plan (projectKey-planKey-jobKey).
	// QUERY KEYS
	MaxResultsKey = "max-result"
)
