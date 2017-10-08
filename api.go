package bamboo

type buildService interface {
	GetBuildResults(max int) Results
	GetProjects() Projects
	GetPlans() Plans
}
