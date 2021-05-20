package main

import (
	"time"

	"github.com/erhemdi/workerpool/worker"
)

func main() {
	workerpool := worker.New(worker.Param{
		NumWorker: 100,
		IsDebug:   true,
	})

	for i := 0; i < 10000; i++ {
		job := worker.Job{
			ID:    i,
			DoJob: mockHitAPI,
		}

		workerpool.SendJob(job)
	}

	workerpool.Wait()
}

func mockHitAPI() {
	time.Sleep(50 * time.Millisecond)
}
