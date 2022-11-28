package httpHandler

import (
	"demo/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router) {
	r.HandleFunc("/jobList", JobListHandler).
		Methods("GET")
	r.HandleFunc("/startJob", StartJobHandler).
		Methods("POST")
	r.HandleFunc("/finishJob/{jobID:[a-zA-Z0-9-]*$}", FinishJobHandler).
		Methods("POST")
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server started")
	log.Fatal(srv.ListenAndServe())
}

func StartJobHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.JobHttpRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	now := time.Now().UTC()
	duration := time.Millisecond * time.Duration(payload.Duration)
	dueTime := now.Add(duration)
	newJob := models.Job{
		JobID:       uuid.NewString(),
		JobName:     payload.JobName,
		DueTime:     dueTime,
		IsCompleted: false,
	}
	models.JobList[newJob.JobID] = &newJob

	resp, err := json.Marshal(newJob)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func FinishJobHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID := vars["jobID"]

	if models.JobList[jobID] == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
	if models.JobList[jobID].IsCompleted {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	if models.JobList[jobID].DueTime.Before(time.Now()) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	models.JobList[jobID].IsCompleted = true
	w.WriteHeader(http.StatusNoContent)
}

func JobListHandler(w http.ResponseWriter, r *http.Request) {
	jobListJson := models.JobListHttpResponse{}
	jobListJson.Data = make([]models.Job, 0)

	for _, job := range models.JobList {
		jobListJson.Data = append(jobListJson.Data, *job)
	}
	resp, err := json.Marshal(jobListJson)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
