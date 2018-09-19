// Better ticker:
//   - .Stop method actually closes tick channel'
//   - .Start method sends first tick right away after call
//   - .Await method for sync usage
package ticker

import (
	"fmt"
	"sync"
	"time"
)

func Every(t time.Duration, action func()) (stop func()) {
	var ticker = NewTicker(t)
	go func() {
		defer ticker.Stop()
		for ticker.Await() {
			action()
		}
	}()
	return ticker.Stop
}

type Ticker struct {
	events chan time.Time
	stop   chan struct{}
	step   time.Duration
	once   *sync.Once
}

func NewTicker(step time.Duration) Ticker {
	return Ticker{
		events: make(chan time.Time),
		stop:   make(chan struct{}),
		once:   &sync.Once{},
		step:   step,
	}
}

func (ticker Ticker) Start() {
	go ticker.tick()
}

func (ticker Ticker) StartAfter(t time.Duration) {
	time.AfterFunc(t, ticker.tick)
}

func (ticker Ticker) Stop() {
	ticker.once.Do(func() {
		close(ticker.stop)
	})
}

func (ticker Ticker) Closed() <-chan struct{} {
	return ticker.stop
}

func (ticker Ticker) Ticks() <-chan time.Time {
	return ticker.events
}

func ExampleTicker_Ticks() {
	var ticker = NewTicker(420 * time.Millisecond)
	ticker.Start()
	for {
		select {
		case <-ticker.Ticks():
			// pass
			// another cases
		}
	}
}

// Await waits for tick, returns true if channel is not closed
func (ticker Ticker) Await() bool {
	var _, ok = <-ticker.events
	return ok
}

func ExampleTicker_Await() {
	var ticker = NewTicker(420 * time.Millisecond)
	ticker.Start()
	defer ticker.Stop() // it's safe to call stop multiple times
	for ticker.Await() {
		fmt.Println("Tick")
		// if something wrong
		ticker.Stop()
	}
}

func (ticker Ticker) tick() {
	select {
	case ticker.events <- time.Now():
		time.AfterFunc(ticker.step, ticker.tick)
	case <-ticker.stop:
		close(ticker.events)
		return
	}
}
