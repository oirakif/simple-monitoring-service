package cronHandler

import (
	"demo/models"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func Init(c *cron.Cron) {
	c.AddFunc("*/5 * * * * *", checkJobs)

	c.Start()
	fmt.Println("cron started")
}

func checkJobs() {
	fmt.Println("checking jobs ...")
	for jobID, job := range models.JobList {
		if job.DueTime.Before(time.Now()) && !job.IsCompleted {
			fmt.Printf("job ID : %s, job name : %s -> failed\n", jobID, job.JobName)
			delete(models.JobList, jobID)
		}
	}
	fmt.Println("check job finished")
}
