package models

import "time"

var JobList = map[string]*Job{}

type Job struct {
	JobID       string    `json:"jobID"`
	JobName     string    `json:"jobName"`
	DueTime     time.Time `json:"dueTime"`
	IsCompleted bool      `json:"isCompleted"`
}

type JobHttpRequest struct {
	JobName  string `json:"jobName"`
	Duration int    `json:"duration"`
}

type JobListHttpResponse struct {
	Data []Job `json:"data"`
}
