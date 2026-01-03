package pipeline

import (
	"fmt"
	"sync/atomic"
)

type Metrics struct {
	success   int64
	failed    int64
	retries   int64
	cancelled int64
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
