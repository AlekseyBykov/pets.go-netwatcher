package models

import (
	"fmt"
	"time"
)

type Job struct {
	URL string
}

type Result struct {
	URL          string
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

func (r Result) Info() string {
	if r.Error != nil {
		return fmt.Sprintf("[ERROR] - [%s] - %s", r.URL, r.Error.Error())
	}
	return fmt.Sprintf("[SUCCESS] - [%s] - Status: %d, Response Time: %s",
		r.URL,
		r.StatusCode,
		r.ResponseTime.String(),
	)
}
