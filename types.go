package goapi

import "time"

const (
	TimeFormat = "2006-01-02T15:04:05"
)

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
	AgentUuid           string               `json:"agent_uuid,omitempty"`
	Name                string               `json:"name"`
	JobStateTransitions []JobStateTransition `json:"job_state_transitions"`
	ScheduledDate       int64                `json:"scheduled_date"`
	PipelineCounter     int                  `json:"pipeline_counter,omitempty"`
	Result              string               `json:"result"`
	State               string               `json:"state"`
	Id                  int                  `json:"id"`
	StageCounter        string               `json:"stage_counter,omitempty"`
	StageName           string               `json:"stage_name,omitempty"`
}

type Stage struct {
	Name            string `json:"name"`
	ApprovedBy      string `json:"approved_by,omitempty"`
	Jobs            []Job  `json:"jobs,omitempty"`
	PipelineCounter int    `json:"pipeline_counter,omitempty"`
	PipelineName    string `json:"pipeline_name,omitempty"`
	ApprovalType    string `json:"approval_type,omitempty"`
	Result          string `json:"result,omitempty"`
	Id              int    `json:"id,omitempty"`
	Counter         string `json:"counter,omitempty"`
}

type Material struct {
	Description string `json:"description"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
}

type Pipeline struct {
	Stages    []Stage    `json:"stages,omitempty"`
	Name      string     `json:"name"`
	Materials []Material `json:"materials,omitempty"`
	Label     string     `json:"label"`
}

type PipelineStatus struct {
	Locked      bool `json:"locked"`
	Paused      bool `json:"paused"`
	Schedulable bool `json:"schedulable"`
}

type PipelineHistory struct {
	Pipelines  []Pipeline `json:"pipelines"`
	Pagination Pagination `json:"pagination"`
}

type PipelineGroup struct {
	Pipelines []Pipeline `json:"pipelines"`
	Name      string     `json:"name"`
}

type Project struct {
	Name                string `xml:"name,attr"`
	Activity            string `xml:"activity,attr"`
	LastBuildStatus     string `xml:"lastBuildStatus,attr"`
	LastBuildLabel      string `xml:"lastBuildLabel,attr"`
	LastBuildTimeString string `xml:"lastBuildTime,attr"`
	WebUrl              string `xml:"webUrl,attr"`
}

func (p Project) LastBuildTime() (time.Time, error) {
	return time.Parse(TimeFormat, p.LastBuildTimeString)
}

type CCTray struct {
	Projects []Project `xml:"Project"`
}
