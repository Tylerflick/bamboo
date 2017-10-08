package bamboo

type ServerSettings struct {
	rootUrl  string
	username string
	password string
}

// Result - Bamboo result
type Result struct {
	BuildNumber    int    `json:"buildNumber"`
	BuildResultKey string `json:"buildResultKey"`
	BuildState     string `json:"buildState"`
	ID             int    `json:"id"`
	Key            string `json:"key"`
	LifeCycleState string `json:"lifeCycleState"`
	Link           struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"link"`
	Number int `json:"number"`
	Plan   struct {
		Enabled bool   `json:"enabled"`
		Key     string `json:"key"`
		Link    struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"link"`
		Name    string `json:"name"`
		PlanKey struct {
			Key string `json:"key"`
		} `json:"planKey"`
		ShortKey  string `json:"shortKey"`
		ShortName string `json:"shortName"`
		Type      string `json:"type"`
	} `json:"plan"`
	PlanResultKey struct {
		EntityKey struct {
			Key string `json:"key"`
		} `json:"entityKey"`
		Key          string `json:"key"`
		ResultNumber int    `json:"resultNumber"`
	} `json:"planResultKey"`
	State string `json:"state"`
}

// Results - Result struct
type Results struct {
	Results struct {
		Expand     string   `json:"expand"`
		MaxResult  int      `json:"max-result"`
		Result     []Result `json:"result"`
		Size       int      `json:"size"`
		StartIndex int      `json:"start-index"`
	} `json:"results"`
}

type ChangeSets struct {
	Author      string `json:"author"`
	ChangesetID string `json:"changesetId"`
}

type BuildPlanResult struct {
	Artifacts struct {
		MaxResult  int `json:"max-result"`
		Size       int `json:"size"`
		StartIndex int `json:"start-index"`
	} `json:"artifacts"`
	BuildCompletedDate       string `json:"buildCompletedDate"`
	BuildCompletedTime       string `json:"buildCompletedTime"`
	BuildDuration            int    `json:"buildDuration"`
	BuildDurationDescription string `json:"buildDurationDescription"`
	BuildDurationInSeconds   int    `json:"buildDurationInSeconds"`
	BuildNumber              int    `json:"buildNumber"`
	BuildReason              string `json:"buildReason"`
	BuildRelativeTime        string `json:"buildRelativeTime"`
	BuildResultKey           string `json:"buildResultKey"`
	BuildStartedTime         string `json:"buildStartedTime"`
	BuildState               string `json:"buildState"`
	BuildTestSummary         string `json:"buildTestSummary"`
	Changes                  struct {
		Change     []ChangeSets `json:"change"`
		Expand     string       `json:"expand"`
		MaxResult  int          `json:"max-result"`
		Size       int          `json:"size"`
		StartIndex int          `json:"start-index"`
	} `json:"changes"`
	Comments struct {
		MaxResult  int `json:"max-result"`
		Size       int `json:"size"`
		StartIndex int `json:"start-index"`
	} `json:"comments"`
	Continuable     bool   `json:"continuable"`
	Expand          string `json:"expand"`
	FailedTestCount int    `json:"failedTestCount"`
	Finished        bool   `json:"finished"`
	ID              int    `json:"id"`
	JiraIssues      struct {
		MaxResult  int `json:"max-result"`
		Size       int `json:"size"`
		StartIndex int `json:"start-index"`
	} `json:"jiraIssues"`
	Key    string `json:"key"`
	Labels struct {
		MaxResult  int `json:"max-result"`
		Size       int `json:"size"`
		StartIndex int `json:"start-index"`
	} `json:"labels"`
	LifeCycleState string `json:"lifeCycleState"`
	Link           struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"link"`
	Metadata struct {
		MaxResult  int `json:"max-result"`
		Size       int `json:"size"`
		StartIndex int `json:"start-index"`
	} `json:"metadata"`
	NotRunYet bool `json:"notRunYet"`
	Number    int  `json:"number"`
	OnceOff   bool `json:"onceOff"`
	Plan      struct {
		Enabled bool   `json:"enabled"`
		Key     string `json:"key"`
		Link    struct {
			Href string `json:"href"`
			Rel  string `json:"rel"`
		} `json:"link"`
		Name    string `json:"name"`
		PlanKey struct {
			Key string `json:"key"`
		} `json:"planKey"`
		ShortKey  string `json:"shortKey"`
		ShortName string `json:"shortName"`
		Type      string `json:"type"`
	} `json:"plan"`
	PlanName      string `json:"planName"`
	PlanResultKey struct {
		EntityKey struct {
			Key string `json:"key"`
		} `json:"entityKey"`
		Key          string `json:"key"`
		ResultNumber int    `json:"resultNumber"`
	} `json:"planResultKey"`
	PrettyBuildCompletedTime string `json:"prettyBuildCompletedTime"`
	PrettyBuildStartedTime   string `json:"prettyBuildStartedTime"`
	ProjectName              string `json:"projectName"`
	QuarantinedTestCount     int    `json:"quarantinedTestCount"`
	ReasonSummary            string `json:"reasonSummary"`
	Restartable              bool   `json:"restartable"`
	SkippedTestCount         int    `json:"skippedTestCount"`
	Stages                   struct {
		MaxResult  int `json:"max-result"`
		Size       int `json:"size"`
		StartIndex int `json:"start-index"`
	} `json:"stages"`
	State               string `json:"state"`
	Successful          bool   `json:"successful"`
	SuccessfulTestCount int    `json:"successfulTestCount"`
	VcsRevisionKey      string `json:"vcsRevisionKey"`
	VcsRevisions        struct {
		MaxResult  int `json:"max-result"`
		Size       int `json:"size"`
		StartIndex int `json:"start-index"`
	} `json:"vcsRevisions"`
}

type AggregatorMessage struct {
	Message   string `json:"message"`
	AlertType string `json:"alert_type"`
}

type Project struct {
	Name string `json:"projects"`
	Key  string `json:"key"`
	Link struct {
		Rel  string `json:"self"`
		Href string `json:"href"`
	}
}

type Projects struct {
	Expand string `json:"expand"`
	Link   struct {
		Rel  string `json:"self"`
		Href string `json:"href"`
	}
	Projects struct {
		Expand     string    `json:"expand"`
		Size       int       `json:"size"`
		MaxResult  int       `json:"max-result"`
		StartIndex int       `json:"start-index"`
		Projects   []Project `json:"project"`
	} `json:"projects"`
}

type Plan struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Link struct {
		Rel  string `json:"self"`
		Href string `json:"href"`
	}
}

type Plans struct {
	Expand string `json:"expand"`
	Link   struct {
		Rel  string `json:"self"`
		Href string `json:"href"`
	}
	Plans struct {
		Expand     string  `json:"expand"`
		Size       int     `json:"size"`
		MaxResult  int     `json:"max-result"`
		StartIndex int     `json:"start-index"`
		Plans      []Plans `json:"plan"`
	} `json:"plans"`
}
