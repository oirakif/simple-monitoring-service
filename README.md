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

## Implementation diagram
<img width="576" alt="image" src="https://user-images.githubusercontent.com/83383024/204227138-afabe8ab-2041-4752-8a3a-f758f563fedb.png">
