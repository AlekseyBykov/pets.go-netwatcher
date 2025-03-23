package main

import (
	"fmt"
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/jobs"
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/pkg/models"
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/pool"
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/processor"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	REQUEST_TIMEOUT = time.Second * 2
	WORKERS_COUNT   = 3
)

func main() {
	resultsCh := make(chan models.Result)
	workerPool := pool.New(WORKERS_COUNT, REQUEST_TIMEOUT, resultsCh)
	workerPool.Init()

	doneCh := make(chan struct{})

	rp := processor.NewResultProcessor()

	go jobs.GenerateJobs(workerPool)
	rp.ProcessResults(resultsCh, doneCh)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGTERM, syscall.SIGINT)
	<-quitCh

	fmt.Println("Shutdown signal received...")
	workerPool.Stop()
	close(resultsCh)

	<-doneCh
	fmt.Println("Graceful shutdown complete")
}
