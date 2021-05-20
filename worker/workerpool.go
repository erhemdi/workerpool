package worker

import (
	"log"
	"sync"
	"time"
)

type Job struct {
	ID    int
	DoJob func()
}

type Param struct {
	NumWorker int
	IsDebug   bool
}

type WorkerPool struct {
	numWorker   int
	chanJob     chan Job
	wg          *sync.WaitGroup
	startTime   time.Time
	isDebugMode bool
}

func New(param Param) *WorkerPool {
	numWorker := param.NumWorker

	chanJob := make(chan Job, numWorker)

	wg := new(sync.WaitGroup)
	wg.Add(numWorker)

	worker := &WorkerPool{
		numWorker:   numWorker,
		chanJob:     chanJob,
		wg:          wg,
		startTime:   time.Now(),
		isDebugMode: param.IsDebug,
	}

	worker.startWorker()

	return worker
}

func (worker *WorkerPool) SendJob(job Job) {
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

	if worker.isDebugMode {
		elapsed := time.Since(worker.startTime).Seconds()
		log.Printf("Process is done - %.3f seconds \n", elapsed)
	}
}

func (worker *WorkerPool) startWorker() {
	for i := 0; i < worker.numWorker; i++ {
		go func() {
			defer worker.wg.Done()

			for job := range worker.chanJob {
				job.DoJob()
			}
		}()
	}

	if worker.isDebugMode {
		log.Printf("%d Worker is waiting for job... \n", worker.numWorker)
	}
}
