package goapi

type Agent struct {
	Sandbox      string   `json:"sandbox"`
	Os           string   `json:"os"`
	Resources    []string `json:"resources,omitempty"`
	Environments []string `json:"environments,omitempty"`
	FreeSpace    string   `json:"free_space"`
	IpAddress    string   `json:"ip_address"`
	AgentName    string   `json:"agent_name"`
	Status       string   `json:"status"`
	Uuid         string   `json:"uuid"`
	BuildLocator string   `json:"build_locator,omitempty"`
}

type Pagination struct {
	Offset   int `json:"offset"`
	Total    int `json:"total"`
	PageSize int `json:"page_size"`
}

type JobStateTransition struct {
	StateChangeTime int64  `json:"state_change_time"`
	Id              int    `json:"id"`
	State           string `json:"scheduled"`
}

type Job struct {
	AgentUuid           string               `json:"agent_uuid"`
	Name                string               `json:"name"`
	JobStateTransitions []JobStateTransition `json:"job_state_transitions"`
	ScheduledDate       int64                `json:"scheduled_date"`
	PipelineCounter     int                  `json:"pipeline_counter"`
	Result              string               `json:"result"`
	State               string               `json:"state"`
	Id                  int                  `json:"id"`
	StageCounter        string               `json:"stage_counter"`
	StageName           string               `json:"stage_name"`
}

type Stage struct {
	Name            string `json:"name"`
	ApprovedBy      string `json:"approved_by"`
	Jobs            []Job  `json:"jobs"`
	PipelineCounter int    `json:"pipeline_counter"`
	PipelineName    string `json:"pipeline_name"`
	ApprovalType    string `json:"approval_type"`
	Result          string `json:"result"`
	Id              int    `json:"id"`
	Counter         string `json:"counter"`
}

type Material struct {
	Description string `json:"description"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
}

type Pipeline struct {
	Stages    []Stage    `json:"stages"`
	Name      string     `json:"name"`
	Materials []Material `json:"materials"`
	Label     string     `json:"label"`
}

type PipelineGroup struct {
	Pipelines []Pipeline `json:"pipelines"`
	Name      string     `json:"name"`
}
