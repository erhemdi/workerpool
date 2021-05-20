package worker

import (
	"log"
	"sync"
	"time"
)

type WorkerPool struct {
	numWorker int
	chanJob   chan func()
	wg        *sync.WaitGroup
	startTime time.Time
}

func New(numWorker int) *WorkerPool {
	chanJob := make(chan func(), numWorker)

	wg := new(sync.WaitGroup)
	wg.Add(numWorker)

	worker := &WorkerPool{
		numWorker: numWorker,
		chanJob:   chanJob,
		wg:        wg,
		startTime: time.Now(),
	}

	worker.startWorker()

	return worker
}

func (worker *WorkerPool) SendJob(job func()) {
	if worker.chanJob == nil {
		log.Println("channel job is nil")
		return
	}

	worker.chanJob <- job
}

func (worker *WorkerPool) Wait() {
	if worker.chanJob == nil {
		log.Println("channel job is nil")
		return
	}

	close(worker.chanJob)
	worker.wg.Wait()

	elapsed := time.Since(worker.startTime).Seconds()
	log.Printf("Process is done - %.3f seconds \n", elapsed)
}

func (worker *WorkerPool) startWorker() {
	for i := 0; i < worker.numWorker; i++ {
		go func() {
			defer worker.wg.Done()

			for job := range worker.chanJob {
				job()
			}
		}()
	}

	log.Printf("%d Worker is waiting for job... \n", worker.numWorker)
}
