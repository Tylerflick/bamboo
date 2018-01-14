package bamboo

const (
	apiURL = "/rest/api"
	// API ENDPOINTS
	latestResults            = "/latest/result"
	projectEndpoint          = "/latest/projects"
	planEndpoint             = "/latest/plan/latest"
	listPlanBranchesEndpoint = "/latest/plan/%s/branch" // buildKey might be top level plan (projectKey-planKey) or job plan (projectKey-planKey-jobKey).
	// QUERY KEYS
	maxResultsKey = "max-result"
)
