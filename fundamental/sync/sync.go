package main


import "sync"

// Counter will increment a number.
type Counter struct {
	mu    sync.Mutex //mutux is a lock. 0 means no lock
	value int
	//dont declare sync.Mutex as syntax sugar in (1) because: embedded typas are exposed as public interface by  implementation. Very dangerous
}

/* 
NewCounter returns a new Counter.

*/
func NewCounter() *Counter {
	return &Counter{}
}

// Inc the count.
func (c *Counter) Inc() {
	c.mu.Lock() // NO c.lock() as syntax sugar (1)
	defer c.mu.Unlock()//releases the lock for next calls
	c.value++
}

// Value returns the current count.
func (c *Counter) Value() int {
	return c.value
}
