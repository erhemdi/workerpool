### Simple Worker Pool

### How to Use
```
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
```

### How to Run
```
go build && ./workerpool.exe
```
