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
		Expand     string `json:"expand"`
		Size       int    `json:"size"`
		MaxResult  int    `json:"max-result"`
		StartIndex int    `json:"start-index"`
		Plans      []Plan `json:"plan"`
	} `json:"plans"`
}

type PlanResult struct {
	Results struct {
		Size       int    `json:"size"`
		Expand     string `json:"expand"`
		StartIndex int    `json:"start-index"`
		MaxResult  int    `json:"max-result"`
		Result     []struct {
			Expand string `json:"expand"`
			Link   struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"link"`
			Plan struct {
				ShortName string `json:"shortName"`
				ShortKey  string `json:"shortKey"`
				Type      string `json:"type"`
				Enabled   bool   `json:"enabled"`
				Link      struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
				} `json:"link"`
				Key     string `json:"key"`
				Name    string `json:"name"`
				PlanKey struct {
					Key string `json:"key"`
				} `json:"planKey"`
			} `json:"plan"`
			PlanName                 string `json:"planName"`
			ProjectName              string `json:"projectName"`
			BuildResultKey           string `json:"buildResultKey"`
			LifeCycleState           string `json:"lifeCycleState"`
			ID                       int    `json:"id"`
			BuildStartedTime         string `json:"buildStartedTime"`
			PrettyBuildStartedTime   string `json:"prettyBuildStartedTime"`
			BuildCompletedTime       string `json:"buildCompletedTime"`
			BuildCompletedDate       string `json:"buildCompletedDate"`
			PrettyBuildCompletedTime string `json:"prettyBuildCompletedTime"`
			BuildDurationInSeconds   int    `json:"buildDurationInSeconds"`
			BuildDuration            int    `json:"buildDuration"`
			BuildDurationDescription string `json:"buildDurationDescription"`
			BuildRelativeTime        string `json:"buildRelativeTime"`
			VcsRevisionKey           string `json:"vcsRevisionKey"`
			VcsRevisions             struct {
				Size       int `json:"size"`
				StartIndex int `json:"start-index"`
				MaxResult  int `json:"max-result"`
			} `json:"vcsRevisions"`
			BuildTestSummary     string `json:"buildTestSummary"`
			SuccessfulTestCount  int    `json:"successfulTestCount"`
			FailedTestCount      int    `json:"failedTestCount"`
			QuarantinedTestCount int    `json:"quarantinedTestCount"`
			SkippedTestCount     int    `json:"skippedTestCount"`
			Continuable          bool   `json:"continuable"`
			OnceOff              bool   `json:"onceOff"`
			Restartable          bool   `json:"restartable"`
			NotRunYet            bool   `json:"notRunYet"`
			Finished             bool   `json:"finished"`
			Successful           bool   `json:"successful"`
			BuildReason          string `json:"buildReason"`
			ReasonSummary        string `json:"reasonSummary"`
			Artifacts            struct {
				Size       int `json:"size"`
				StartIndex int `json:"start-index"`
				MaxResult  int `json:"max-result"`
			} `json:"artifacts"`
			Comments struct {
				Size       int `json:"size"`
				StartIndex int `json:"start-index"`
				MaxResult  int `json:"max-result"`
			} `json:"comments"`
			Labels struct {
				Size       int `json:"size"`
				StartIndex int `json:"start-index"`
				MaxResult  int `json:"max-result"`
			} `json:"labels"`
			JiraIssues struct {
				Size       int `json:"size"`
				StartIndex int `json:"start-index"`
				MaxResult  int `json:"max-result"`
			} `json:"jiraIssues"`
			Stages struct {
				Size       int `json:"size"`
				StartIndex int `json:"start-index"`
				MaxResult  int `json:"max-result"`
			} `json:"stages"`
			Key           string `json:"key"`
			PlanResultKey struct {
				Key       string `json:"key"`
				EntityKey struct {
					Key string `json:"key"`
				} `json:"entityKey"`
				ResultNumber int `json:"resultNumber"`
			} `json:"planResultKey"`
			State       string `json:"state"`
			BuildState  string `json:"buildState"`
			Number      int    `json:"number"`
			BuildNumber int    `json:"buildNumber"`
		} `json:"result"`
	} `json:"results"`
	Expand string `json:"expand"`
	Link   struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"link"`
}

type PlanDetails struct {
	Key  string `json:"key"`
	Link struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"link"`
	Name        string `json:"name"`
	ProjectKey  string `json:"projectKey"`
	ProjectName string `json:"projectName"`
	ShortKey    string `json:"shortKey"`
	ShortName   string `json:"shortName"`
}

type PlanBranches struct {
	Expand string `json:"expand"`
	Link   struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"link"`
	Branches struct {
		Size       int    `json:"size"`
		Expand     string `json:"expand"`
		StartIndex int    `json:"start-index"`
		MaxResult  int    `json:"max-result"`
		Branch     []struct {
			Description string `json:"description"`
			ShortName   string `json:"shortName"`
			ShortKey    string `json:"shortKey"`
			Enabled     bool   `json:"enabled"`
			Link        struct {
				Href string `json:"href"`
				Rel  string `json:"rel"`
			} `json:"link"`
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"branches"`
}
