package pipeline

import (
	"context"
	"sync"
)

type WorkerPool struct {
	workers int
	limiter chan struct{}
	metrics *Metrics
}

func NewWorkerPool(workers, maxInflight int, m *Metrics) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		limiter: make(chan struct{}, maxInflight),
		metrics: m,
	}
}

func (p *WorkerPool) Start(ctx context.Context, jobs <-chan ImageJob, wg *sync.WaitGroup) {
	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		go p.worker(ctx, jobs, wg)
	}
}

func (p *WorkerPool) worker(ctx context.Context, jobs <-chan ImageJob, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			p.limiter <- struct{}{}
			p.process(ctx, job)
			<-p.limiter
		}
	}
}

func (p *WorkerPool) process(ctx context.Context, job ImageJob) {
	for job.AttemptsLeft > 0 {
		select {
		case <-ctx.Done():
			p.metrics.Cancelled()
			return
		default:
		}

		err := ProcessImage(job, func(size int) {
			// optional: per-thumbnail progress
		})
		if err == nil {
			p.metrics.Success()
			return
		}

		job.AttemptsLeft--
		p.metrics.Retry()
	}

	p.metrics.Failed()
}
