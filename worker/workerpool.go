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
	Name      string
	NumWorker int
	IsDebug   bool
}

type workerPool struct {
	name        string
	numWorker   int
	chanJob     chan Job
	wg          *sync.WaitGroup
	startTime   time.Time
	isDebugMode bool
}

func New(param Param) *workerPool {
	numWorker := param.NumWorker

	chanJob := make(chan Job, numWorker)

	wg := new(sync.WaitGroup)
	wg.Add(numWorker)

	worker := &workerPool{
		name:        param.Name,
		numWorker:   numWorker,
		chanJob:     chanJob,
		wg:          wg,
		startTime:   time.Now(),
		isDebugMode: param.IsDebug,
	}

	return worker
}

func (worker *workerPool) SendJob(job Job) {
	if worker.chanJob == nil {
		log.Println("channel job is nil")
		return
	}

	worker.chanJob <- job
}

func (worker *workerPool) Wait() {
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

func (worker *workerPool) Run() {
	for i := 0; i < worker.numWorker; i++ {
		go func() {
			defer worker.wg.Done()

			for job := range worker.chanJob {
				job.DoJob()
			}
		}()
	}

	if worker.isDebugMode {
		log.Printf("%d Worker %s is waiting for job... \n",
			worker.numWorker, worker.name)
	}
}
