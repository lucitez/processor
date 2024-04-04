package util

import "sync"

// SafeMap is safe to use concurrently.
type SafeMap[K comparable, V any] interface {
	Store(key K, val V) (ok bool)
	Get(key K) V
	Map() map[K]V
	NumKeys() int
}

type SafeMapBoolean[K comparable] struct {
	mu   sync.Mutex
	SMap map[K]bool
}

func (b *SafeMapBoolean[K]) Store(key K, val bool) (ok bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.SMap[key] {
		return false
	}

	b.SMap[key] = val
	return true
}

func (b *SafeMapBoolean[K]) Get(key K) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.SMap[key]
}

func (b *SafeMapBoolean[K]) Map() map[K]bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.SMap
}

func (b *SafeMapBoolean[K]) NumKeys() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.SMap)
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

func (c *SafeMapCounter[K]) Get(key K) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.SMap[key]
}

func (c *SafeMapCounter[K]) Map() map[K]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.SMap
}

func (c *SafeMapCounter[K]) NumKeys() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.SMap)
}
