package session

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/mamaart/freestyle-timer/internal/timer"
	"github.com/mamaart/freestyle-timer/models"
)

type athlete struct {
	Name  string
	Timer *timer.Timer
}

func newAthlete(name string, timer *timer.Timer) *athlete {
	return &athlete{
		Name:  name,
		Timer: timer,
	}
}

func (a *athlete) String() string {
	return fmt.Sprintf("%s ::: %s", a.Timer.String(), a.Name)
}

type Session struct {
	mu          sync.Mutex
	title       string
	athleteOne  *athlete
	athleteTwo  *athlete
	activeTimer *timer.Timer
}

func New(opt models.Opt, bc1, bc2 chan<- time.Duration) *Session {
	return &Session{
		title:      opt.SessionTitle,
		athleteOne: newAthlete(opt.AthleteOneName, timer.New(bc1)),
		athleteTwo: newAthlete(opt.AthleteTwoName, timer.New(bc2)),
	}
}

func (s *Session) withTimer(i int, fn func(t *timer.Timer) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch i {
	case 1:
		return fn(s.athleteOne.Timer)
	case 2:
		return fn(s.athleteTwo.Timer)
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
		"%s\n%s\n",
		s.athleteOne.String(),
		s.athleteTwo.String(),
	)
}
