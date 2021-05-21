package main

import (
	"time"

	"github.com/erhemdi/workerpool/worker"
)

func main() {
	workerPool := worker.New(worker.Param{
		Name:      "mock_hit_API",
		NumWorker: 100,
		IsDebug:   true,
	})
	workerPool.Run()

	for i := 0; i < 10000; i++ {
		job := worker.Job{
			ID:    i,
			DoJob: mockHitAPI,
		}

		workerPool.SendJob(job)
	}

	workerPool.Wait()
}

func mockHitAPI() {
	time.Sleep(50 * time.Millisecond)
}
