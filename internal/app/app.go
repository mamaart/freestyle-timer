package app

import (
	"errors"
	"sync"
	"time"

	"github.com/mamaart/freestyle-timer/internal/session"
	"github.com/mamaart/freestyle-timer/models"
)

type App struct {
	mu           sync.Mutex
	session      *session.Session
	broadcastOne chan time.Duration
	broadcastTwo chan time.Duration
	regOne       chan chan<- time.Duration
	regTwo       chan chan<- time.Duration
}

func New() *App {
	broadcastOne := make(chan time.Duration)
	regOne := make(chan chan<- time.Duration)
	go registerAndBroadcast(broadcastOne, regOne)

	broadcastTwo := make(chan time.Duration)
	regTwo := make(chan chan<- time.Duration)
	go registerAndBroadcast(broadcastTwo, regTwo)

	return &App{
		broadcastOne: broadcastOne,
		regOne:       regOne,

		broadcastTwo: broadcastTwo,
		regTwo:       regTwo,
	}
}

func registerAndBroadcast(broadcast <-chan time.Duration, register <-chan chan<- time.Duration) {
	listeners := make(map[chan<- time.Duration]struct{})
	for {
		select {
		case l := <-register:
			listeners[l] = struct{}{}
		case d := <-broadcast:
			for l := range listeners {
				select {
				case l <- d:
				default:
					delete(listeners, l)
				}
			}
		}
	}
}

func (a *App) NewSession(opt models.Opt) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.session != nil {
		return errors.New("current session not finnished")
	}

	a.session = session.New(opt, a.broadcastOne, a.broadcastTwo)

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

func (a *App) Listen(i int) (chan time.Duration, error) {
	if a.session == nil {
		return nil, errors.New("no current session exist")
	}

	ch := make(chan time.Duration)

	switch i {
	case 1:
		a.regOne <- ch
	case 2:
		a.regTwo <- ch
	default:
		return nil, errors.New("invalid athlete number")
	}

	return ch, nil
}
