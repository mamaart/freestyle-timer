package session

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/mamaart/freestyle-timer/internal/timer"
)

type Session struct {
	mu          sync.Mutex
	athleteOne  *timer.Timer
	athleteTwo  *timer.Timer
	activeTimer *timer.Timer
}

func New(d time.Duration) *Session {
	return &Session{
		athleteOne: timer.New(d),
		athleteTwo: timer.New(d),
	}
}

func (s *Session) withTimer(i int, fn func(t *timer.Timer) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch i {
	case 1:
		return fn(s.athleteOne)
	case 2:
		return fn(s.athleteTwo)
	}

	return errors.New("invalid timer number")
}

func (s *Session) Start(i int) error {
	if s.activeTimer != nil {
		return errors.New("another timer is running")
	}

	return s.withTimer(i, func(t *timer.Timer) error {
		s.activeTimer = t
		s.activeTimer.Start()
		return nil
	})
}

func (s *Session) Pause(i int) error {
	return s.withTimer(i, func(t *timer.Timer) error {
		if s.activeTimer != t {
			return errors.New("cannot pause timer which is not active")
		}
		s.activeTimer.Pause()
		s.activeTimer = nil
		return nil
	})
}

func (s *Session) String() string {
	return fmt.Sprintf(
		"Athlete One: %s\nAthlete Two: %s\n",
		s.athleteOne.String(),
		s.athleteTwo.String(),
	)
}
