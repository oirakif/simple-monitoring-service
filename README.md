# Simple Monitoring Service

## Run the service
```
go run main.go
```

## HTTP API Modules

### Get Job List
```
curl --location --request GET 'http://localhost:8000/jobList'
```

### Start a job
This module will return a job ID which needed as a parameter in `finishJob` endpoint
```
curl --location --request POST 'http://localhost:8000/startJob' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jobName":<jobName>,
    "duration":<duration(milisec)>
}'
```

### Finish a job
```
curl --location --request POST 'http://localhost:8000/finishJob/:jobID' \
--header 'Content-Type: application/json'
```