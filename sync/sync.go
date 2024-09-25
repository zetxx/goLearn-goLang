package main

import "sync"

type Counter struct {
	i  int64
	mu sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.i = c.i + 1
}

func (c *Counter) Value() int64 {
	return c.i
}
