# goapi
go client library for ThoughtWorks Go continuous delivery server


# API

## Pipeline Group

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List Pipeline Groups](http://www.go.cd/documentation/user/current/api/pipeline_group_api.html) | /go/api/config/pipeline_groups | GET | TBD | 


## Pipeline

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [Schedule Pipeline](http://www.go.cd/documentation/user/current/api/pipeline_api.html) | /go/api/pipelines/\[pipeline\]/schedule | POST | TBD | 

## Stages

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [Cancel](http://www.go.cd/documentation/user/current/api/stages_api.html) | /go/api/stages/\[pipeline\]/\[stage\]/cancel | POST | TBD | 
| [View](http://www.go.cd/documentation/user/current/api/stages_api.html) | /go/api/stages/\[pipeline\]/\[stage\]/instance/\[pipeline-counter\]/\[stage-counter\] | GET | TBD | 
| [History](http://www.go.cd/documentation/user/current/api/stages_api.html) | /go/api/stages/\[pipeline\]/\[stage\]/history/\[offset\] | GET | TBD | 

## Jobs

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [Scheduled](http://www.go.cd/documentation/user/current/api/jobs_api.html) | /go/api/jobs/scheduled.xml | GET | TBD | 
| [History](http://www.go.cd/documentation/user/current/api/jobs_api.html) | /go/api/\[pipeline\]/\[stage\]/\[job\]/history/\[offset\] | GET | TBD | 

## Agent

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List](http://www.go.cd/documentation/user/current/api/agent_api.html) | /go/api/agents | GET | TBD | 
| [Enable](http://www.go.cd/documentation/user/current/api/agent_api.html) | /go/api/agents/\[agent-uuid\]/enable | POST | TBD | 
| [Disable](http://www.go.cd/documentation/user/current/api/agent_api.html) | /go/api/agents/\[agent-uuid\]/disable | POST | TBD | 
| [Delete](http://www.go.cd/documentation/user/current/api/agent_api.html) | /go/api/agents/\[agent-uuid\]/delete | POST | TBD | 
| [History](http://www.go.cd/documentation/user/current/api/agent_api.html) | /go/api/agents/\[agent-uuid\]/job_run_history/\[offset\] | POST | TBD | 

## Materials

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List](http://www.go.cd/documentation/user/current/api/materials_api.html) | /go/api/config/materials | GET | TBD | 
| [Subversion](http://www.go.cd/documentation/user/current/api/materials_api.html) | /go/api/material/notify/svn | POST | TBD | 
| [Git](http://www.go.cd/documentation/user/current/api/materials_api.html) | /go/api/material/notify/git | POST | TBD | 
| [Modifications](http://www.go.cd/documentation/user/current/api/materials_api.html) | /go/api/materials/\[fingerprint\]/modifications/\[offset\] | POST | TBD | 

## Configuration

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List](http://www.go.cd/documentation/user/current/api/configuration_api.html) | /go/api/config/revisions | GET | TBD | 
| [Subversion](http://www.go.cd/documentation/user/current/api/configuration_api.html) | /go/api/config/diff/\[from-SHA\]/\[to-SHA\] | GET | TBD | 

## Artifacts

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List](http://www.go.cd/documentation/user/current/api/artifacts_api.html) | /go/api/files/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\].json | GET | TBD | 
| [Show](http://www.go.cd/documentation/user/current/api/artifacts_api.html) | /go/api/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\]/\[filename\] | GET | TBD | 
| [Show](http://www.go.cd/documentation/user/current/api/artifacts_api.html) | /go/api/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\]/\[pathname\]/\[filename\] | GET | TBD | 
| [Show](http://www.go.cd/documentation/user/current/api/artifacts_api.html) | /go/api/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\]/\[pathname\].zip | GET | TBD | 
| [Create](http://www.go.cd/documentation/user/current/api/artifacts_api.html) | /go/api/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\]/\[filename\] | POST | TBD | 
| [Update](http://www.go.cd/documentation/user/current/api/artifacts_api.html) | /go/api/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\]/\[filename\] | PUT | TBD | 

## Users

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [Delete](http://www.go.cd/documentation/user/current/api/users_api.html) | /go/api/users/\[user_name\] | DELETE | TBD | 

## Backup

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [Trigger Backup](http://www.go.cd/documentation/user/current/api/backup_api.html) | /go/api/admin/start_backup | DELETE | TBD | 

## Properties

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List](http://www.go.cd/documentation/user/current/api/properties_api.html) | /go/api/properties/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\] | GET | TBD | 
| [Show](http://www.go.cd/documentation/user/current/api/properties_api.html) | /go/api/properties/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\]/\[property-name\] | GET | TBD | 
| [Create](http://www.go.cd/documentation/user/current/api/properties_api.html) | /go/api/properties/\[pipeline\]/\[pipeline-counter\]/\[stage\]/\[stage-counter\]/\[job\]/\[property-name\] | POST | TBD | 
| [Search](http://www.go.cd/documentation/user/current/api/properties_api.html) | /go/api/properties/search | POST | TBD | 

## Feeds

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List All Pipelines](http://www.go.cd/documentation/user/current/api/feeds_api.html) | /go/api/pipelines.xml | GET | TBD | 
| [Pipeline Feed](http://www.go.cd/documentation/user/current/api/feeds_api.html) | /go/api/\[pipeline\]/\[pipeline-id\].xml | GET | TBD | 
| [All Stages for a Pipeline Feed](http://www.go.cd/documentation/user/current/api/feeds_api.html) | /go/api/\[pipeline\]/stages.xml | GET | TBD | 
| [Stage XML](http://www.go.cd/documentation/user/current/api/feeds_api.html) | /go/api/\[stages\]/\[stage-id\].xml | GET | TBD | 
| [Stage XML](http://www.go.cd/documentation/user/current/api/feeds_api.html) | /go/api/\[stage\]/\[stage-counter\].xml | GET | TBD | 
| [Stage XML](http://www.go.cd/documentation/user/current/api/feeds_api.html) | /go/api/\[pipeline\]/\[pipeline-label\]/\[stage\]/\[stage-counter\].xml | GET | TBD | 
| [Job XML](http://www.go.cd/documentation/user/current/api/feeds_api.html) | /go/api/jobs/\[job-id\].xml | GET | TBD | 

## Command Repo

| API | Endpoint | Method | Status |
| --- | ------ | ------ | ------ | 
| [List All Pipelines](http://www.go.cd/documentation/user/current/api/command_repo_api.html) | /go/api/admin/command-repo-cache/reload | POST | TBD | 
