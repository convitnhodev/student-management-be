package asynJob

import (
	"context"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	GetState() JobState
}

const (
	defaultMaxTimeout    = 10 * time.Second
	defaultMaxRetryCount = 3
)

var (
	defaultRetrytime = []time.Duration{time.Second, 5 * time.Second, 10 * time.Second}
)

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateTimeout
	StateFailed
	StateCompleted
	StateRetryFailed
)

func (js *JobState) String() string {
	return []string{"Init", "Running", "Timeout", "Failed", "Completed", "RetryFailed"}[*js]
}

type jobHandler = func(ctx context.Context) error

type job struct {
	config     jobConfig
	handler    jobHandler
	state      JobState
	indexRetry int
	stopChan   chan bool
}

func NewJob(handler jobHandler) *job {
	return &job{
		handler:    handler,
		state:      StateInit,
		indexRetry: -1,
		config: jobConfig{
			defaultMaxTimeout,
			defaultRetrytime,
		},
		stopChan: make(chan bool),
	}
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning
	err := j.handler(ctx)
	if err == nil {
		j.state = StateCompleted
		return nil
	} else {
		j.state = StateFailed
		return err
	}
	return err
}

func (j *job) Retry(ctx context.Context) error {
	j.state = StateRunning
	j.indexRetry += 1

	err := j.handler(ctx)

	if err != nil {
		if j.indexRetry == len(j.config.Retries)-1 {
			j.state = StateRetryFailed
			return err
		}
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil
}

func (j *job) GetState() JobState {
	return j.state
}
