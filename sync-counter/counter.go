package synccounter

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Val() int {
	return c.value
}