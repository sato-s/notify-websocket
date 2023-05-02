package main

import (
	"log"

	"github.com/kataras/neffos"
)

type Room struct {
	name         string
	ws           *neffos.Server
	sessionTimer *SessionTimer
}

func NewRoom(name string, ws *neffos.Server) *Room {
	r := &Room{
		name:         name,
		ws:           ws,
		sessionTimer: NewSessionTimer(),
	}
	go r.run()
	return r
}

func (r *Room) run() {
	log.Printf("Room `%s` is running\n", r.name)
	r.sessionTimer.Start(10)
	for {
		select {
		case tickString := <-r.sessionTimer.C:
			log.Println("tick")
			m := neffos.Message{
				Namespace: "default",
				Event:     "tick",
				Room:      r.name,
				Body:      []byte(tickString),
			}
			r.ws.Broadcast(nil, m)
		}
	}
}
