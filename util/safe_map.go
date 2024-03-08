package util

import "sync"

// SafeMap is safe to use concurrently.
type SafeMap[K comparable, V any] interface {
	Store(key K, val V)
	Access(key K) V
	Get() map[K]V
}

type SafeMapCounter[K comparable] struct {
	mu   sync.Mutex
	SMap map[K]int
}

func (c *SafeMapCounter[K]) Store(key K, val int) {
	c.mu.Lock()
	c.SMap[key] += val
	c.mu.Unlock()
}

func (c *SafeMapCounter[K]) Access(key K) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.SMap[key]
}

func (c *SafeMapCounter[K]) Get() map[K]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.SMap
}
