package asynJob

import (
	"context"
	"sync"
)

type groupJob struct {
	isConcurrent bool
	jobs         []Job
	wg           *sync.WaitGroup
}

func NewGroupJob(isConcurrent bool, jobs ...Job) *groupJob {
	g := groupJob{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           new(sync.WaitGroup),
	}
	return &g
}

func (g *groupJob) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))
	errChan := make(chan error, len(g.jobs))
	for i, _ := range g.jobs {
		if g.isConcurrent == true {
			go func(job Job) {
				errChan <- g.runJob(ctx, job)
			}(g.jobs[i])

			continue
		}

		errChan <- g.runJob(ctx, g.jobs[i])
		g.wg.Done()
	}

	var err error
	for i := 0; i < len(g.jobs); i++ {
		v := <-errChan
		if v != nil {
			err = v
		}
	}

	g.wg.Wait()
	return err
}

func (g *groupJob) runJob(ctx context.Context, job Job) error {
	err := job.Execute(ctx)
	if err == nil {
		return nil
	}

	for {
		err = job.Retry(ctx)

		if job.GetState() == StateRetryFailed {
			return err
		}

		if err == nil {
			return nil
		}
	}
}
