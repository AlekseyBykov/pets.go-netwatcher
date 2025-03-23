package pool

import (
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/pkg/models"
	"log"
	"sync"
	"time"
)

type Pool struct {
	worker       *worker
	workersCount int

	jobsCh    chan models.Job
	resultsCh chan models.Result

	wg      *sync.WaitGroup
	stopped bool
}

func New(workersCount int, timeout time.Duration, resultsCh chan models.Result) *Pool {
	return &Pool{
		worker:       newWorker(timeout),
		workersCount: workersCount,
		jobsCh:       make(chan models.Job),
		resultsCh:    resultsCh,
		wg:           new(sync.WaitGroup),
	}
}

func (p *Pool) Init() {
	for i := 0; i < p.workersCount; i++ {
		go p.initWorker(i)
	}
}

func (p *Pool) Push(j models.Job) {
	if p.stopped {
		return
	}

	p.jobsCh <- j
	p.wg.Add(1)
}

func (p *Pool) Stop() {
	p.stopped = true
	close(p.jobsCh)
	p.wg.Wait()
}

func (p *Pool) initWorker(id int) {
	// blocking read
	for job := range p.jobsCh {
		time.Sleep(time.Second)

		p.resultsCh <- p.worker.process(job)
		p.wg.Done()
	}

	log.Printf("[worker ID %d] finished proccesing", id)
}
