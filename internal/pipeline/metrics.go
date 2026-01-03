package pipeline

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type Metrics struct {
	success   int64
	failed    int64
	retries   int64
	cancelled int64
	Total     int64
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) Success()   { atomic.AddInt64(&m.success, 1) }
func (m *Metrics) Failed()    { atomic.AddInt64(&m.failed, 1) }
func (m *Metrics) Retry()     { atomic.AddInt64(&m.retries, 1) }
func (m *Metrics) Cancelled() { atomic.AddInt64(&m.cancelled, 1) }

func (m *Metrics) Print() {
	fmt.Printf(`
Processed:
  Success:   %d
  Failed:    %d
  Retries:   %d
  Cancelled: %d
`,
		m.success, m.failed, m.retries, m.cancelled)
}

func (p *WorkerPool) StartProgress(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				fmt.Printf(
					"\rProcessed %d/%d | Success:%d Retry:%d Failed:%d",
					atomic.LoadInt64(&p.metrics.success)+atomic.LoadInt64(&p.metrics.failed),
					atomic.LoadInt64(&p.metrics.Total),
					atomic.LoadInt64(&p.metrics.success),
					atomic.LoadInt64(&p.metrics.retries),
					atomic.LoadInt64(&p.metrics.failed),
				)
			}
		}
	}()
}
