package count

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// TTL 一个周期
const TTL = 5

// ErrTooHigh define
var ErrTooHigh = errors.New("too high")

// Bucket define
type Bucket struct {
	time int64
	data int64
}

// NewBucket define
func NewBucket() *Bucket {
	bucket := &Bucket{
		data: 0,
		time: 1,
	}
	return bucket
}

// Counter 计数器
type Counter struct {
	buckets    [5]*Bucket
	breakValue int64
	mu         *sync.Mutex
}

// Option wrapper
type Option func(*Counter)

// BreakValue will change break value
func BreakValue(value int64) Option {
	return func(c *Counter) {
		c.breakValue = value
	}
}

// NewCounter do init
func NewCounter(options ...func(*Counter)) (*Counter, error) {
	counter := &Counter{
		breakValue: 20,
	}

	for index := range counter.buckets {
		i := index
		counter.buckets[i] = NewBucket()
	}

	for _, option := range options {
		option(counter)
	}

	return counter, nil
}

func (c *Counter) flush() {
	for index := range c.buckets {
		i := index
		c.buckets[i] = NewBucket()
	}
}

// Incr define
func (c *Counter) Incr() error {
	currentTime := time.Now().Unix()
	target := c.buckets[currentTime%int64(len(c.buckets))]
	if target.time == currentTime {
		atomic.AddInt64(&target.time, 1)
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if currentTime-target.time == TTL {
		target.time = currentTime
		target.data = 1
		return nil
	}
	for index := range c.buckets {
		i := index
		c.buckets[i].data = 1
		c.buckets[i].time = currentTime
	}
	return nil
}
