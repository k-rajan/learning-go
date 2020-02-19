package sync

import "sync"

// Counter type
type Counter struct {
	mu    sync.Mutex
	count int
}

// NewCounter empty
func NewCounter() *Counter {
	return &Counter{}
}

// Inc counts
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

// Value of counts
func (c *Counter) Value() int {
	return c.count
}
