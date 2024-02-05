package timer

import (
	"fmt"
	"log"
	"time"
)

type Timer struct {
	remaining time.Duration
	ticker    *time.Ticker
	control   chan bool
	isRunning bool
	done      chan bool
}

func New(d time.Duration) *Timer {
	t := &Timer{
		remaining: d,
		control:   make(chan bool),
		isRunning: false,
		done:      make(chan bool),
		ticker:    time.NewTicker(time.Second),
	}

	t.ticker.Stop()

	go t.run()

	return t
}

func (t *Timer) String() string {
	m := int(t.remaining.Minutes())
	s := int(t.remaining.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", m, s)
}

func (t *Timer) Start() {
	t.control <- true
}

func (t *Timer) Pause() {
	t.control <- false
}

func (t *Timer) Await() {
	<-t.done
}

func (t *Timer) tick() bool {
	if t.remaining > 0 {
		t.remaining -= time.Second
		return false
	}
	t.ticker.Stop()
	t.isRunning = false
	select {
	case t.done <- true:
	default:
	}
	return true
}

func (t *Timer) run() {
	for {
		select {
		case <-t.ticker.C:
			if done := t.tick(); done {
				return
			}
		case start := <-t.control:
			if t.isRunning && !start {
				t.ticker.Stop()
				t.isRunning = false
			} else if !t.isRunning && start {
				t.ticker.Reset(time.Second)
				t.isRunning = true
			} else {
				log.Println("WARNING:", "invalid timer command")
			}
		}
	}
}
