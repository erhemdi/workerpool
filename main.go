package main

import (
	"time"

	"github.com/erhemdi/workerpool/worker"
)

func main() {
	workerpool := worker.New(100)

	for i := 0; i < 10000; i++ {
		workerpool.SendJob(func() {
			mockHitAPI()
		})
	}

	workerpool.Wait()
}

func mockHitAPI() {
	time.Sleep(50 * time.Millisecond)
}
