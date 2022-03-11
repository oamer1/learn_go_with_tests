package sync

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	//  All the other goroutines will have to wait for it to be Unlocked before getting access.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
