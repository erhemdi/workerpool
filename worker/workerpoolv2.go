package worker

import (
	"log"
	"time"
)

type workerPoolV2 struct {
	name        string
	numWorker   int
	chanJob     chan Job
	startTime   time.Time
	isDebugMode bool
	quit        chan bool
}

func NewV2(param Param) workerPoolV2 {
	numWorker := param.NumWorker

	chanJob := make(chan Job, numWorker)

	worker := workerPoolV2{
		name:        param.Name,
		numWorker:   numWorker,
		chanJob:     chanJob,
		startTime:   time.Now(),
		isDebugMode: param.IsDebug,
		quit:        make(chan bool),
	}

	return worker
}

func (worker *workerPoolV2) SendJob(job Job) {
	if worker.chanJob == nil {
		log.Println("channel job is nil")
		return
	}

	worker.chanJob <- job
}

func (worker *workerPoolV2) Run() {
	for i := 0; i < worker.numWorker; i++ {
		go func() {
			for {
				select {
				case job := <-worker.chanJob:
					job.DoJob()
				case <-worker.quit:
					return
				}
			}
		}()
	}

	if worker.isDebugMode {
		log.Printf("%d Worker %s is waiting for job... \n",
			worker.numWorker, worker.name)
	}
}

func (worker *workerPoolV2) Stop() {
	if worker.quit == nil {
		log.Printf("channel quit is nil")
		return
	}

	close(worker.quit)
}
