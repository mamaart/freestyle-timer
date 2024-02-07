package app

import (
	"errors"
	"sync"
	"time"

	"github.com/mamaart/freestyle-timer/internal/session"
	"github.com/mamaart/freestyle-timer/models"
)

type App struct {
	mu      sync.Mutex
	session *session.Session
}

func New() *App {
	return &App{}
}

func (a *App) NewSession(opt models.Opt) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.session != nil {
		return errors.New("current session not finnished")
	}

	a.session = session.New(opt)

	return nil
}

func (a *App) DestroyCurrentSession() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.session == nil {
		return errors.New("no current session exist")
	}

	a.session = nil

	return nil
}

func (a *App) StartTimer(i int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.session == nil {
		return errors.New("no current session exist")
	}

	return a.session.Start(i)
}

func (a *App) PauseTimer(i int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.session == nil {
		return errors.New("no current session exist")
	}

	return a.session.Pause(i)
}

func (a *App) GetState() (string, error) {
	if a.session == nil {
		return "", errors.New("no current session exist")
	}

	return a.session.String(), nil
}

func (a *App) Listen() (chan string, error) {
	if a.session == nil {
		return nil, errors.New("no current session exist")
	}

	ch := make(chan string)

	go func() {
		for {
			if a.session == nil {
				close(ch)
				return
			}
			ch <- a.session.String()
			time.Sleep(time.Second)
		}
	}()

	return ch, nil
}
