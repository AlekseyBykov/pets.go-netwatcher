package jobs

import (
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/pkg/models"
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/pool"
	"time"
)

const INTERVAL = time.Second * 10

var urls = []string{
	"https://github.com",
	"https://google.com",
	"https://golang.org",
	"https://api.coincap.io/v2/assets",
}

func GenerateJobs(p *pool.Pool) {
	go func() {
		for {
			for _, url := range urls {
				// workers wake up
				p.Push(models.Job{URL: url})
			}
			time.Sleep(INTERVAL)
		}
	}()
}
