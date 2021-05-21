package main

import (
	"time"

	"github.com/erhemdi/workerpool/worker"
)

func main() {
	// workerV1()
	// workerV2()
}

func workerV1() {
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

func workerV2() {
	workerPool := worker.NewV2(worker.Param{
		Name:      "mock_hit_API",
		NumWorker: 100,
		IsDebug:   true,
	})

	workerPool.Run()

	for i := 0; i < 1000; i++ {
		job := worker.Job{
			ID:    i,
			DoJob: mockHitAPI,
		}

		workerPool.SendJob(job)
	}
}

func mockHitAPI() {
	time.Sleep(50 * time.Millisecond)
}
