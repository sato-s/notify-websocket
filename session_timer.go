package main

import (
	"fmt"
	"time"
)

const INTERVAL = 1 * time.Second

type SessionTimer struct {
	C           chan string
	tick        chan struct{}
	stop        chan struct{}
	limitTick   int
	currentTick int
	ticker      *time.Ticker
}

func NewSessionTimer() *SessionTimer {
	st := &SessionTimer{
		C:    make(chan string),
		tick: make(chan struct{}),
		stop: make(chan struct{}),
	}
	go st.run()
	return st
}

func (st *SessionTimer) Start(limitTick int) {
	st.limitTick = limitTick
	st.currentTick = 0
	st.ticker = time.NewTicker(INTERVAL)

	go func() {
		for {
			select {
			case <-st.ticker.C:
				st.currentTick++
				if st.currentTick >= st.limitTick {
					st.stop <- struct{}{}
				} else {
					st.tick <- struct{}{}
				}
			}
		}
	}()
}

func (st *SessionTimer) Stop() {
}

func (st *SessionTimer) Reset() {
}

func (st *SessionTimer) run() {
	for {
		select {
		case <-st.tick:
			st.C <- fmt.Sprintf("tick:%d", st.currentTick)
		case <-st.stop:
			return
		}
	}
}
