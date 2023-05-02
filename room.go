package main

import (
	"log"
	"time"

	"github.com/kataras/neffos"
)

type Room struct {
	name   string
	ws     *neffos.Server
	ticker *time.Ticker
}

func NewRoom(name string, ws *neffos.Server) *Room {
	r := &Room{
		name:   name,
		ws:     ws,
		ticker: time.NewTicker(1 * time.Second),
	}
	go r.run()
	return r
}

func (r *Room) run() {
	log.Printf("Room `%s` is running\n", r.name)
	go func() {
		for {
			select {
			case <-r.ticker.C:
				log.Println("tick")
				m := neffos.Message{
					Namespace: "default",
					Event:     "tick",
					Room:      r.name,
					Body:      []byte("tick!"),
				}
				r.ws.Broadcast(nil, m)
			}
		}
	}()
}
