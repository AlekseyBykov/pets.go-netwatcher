package processor

import (
	"fmt"
	"github.com/AlekseyBykov/pets.go-netwatcher/internal/pkg/models"
)

type ResultProcessor struct{}

func NewResultProcessor() *ResultProcessor {
	return &ResultProcessor{}
}

func (rp *ResultProcessor) ProcessResults(resultsCh chan models.Result, doneCh chan struct{}) {
	go func() {
		defer close(doneCh)

		// blocking read
		for result := range resultsCh {
			fmt.Println(result.Info())
		}
	}()
}
