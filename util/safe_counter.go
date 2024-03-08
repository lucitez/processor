package util

import "sync"

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu       sync.Mutex
	numEvens int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.numEvens++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.numEvens
}
