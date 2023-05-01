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
		ticker: time.NewTicker(1 * time.Second),
	}
	r.run()
	return r
}

func (r *Room) run() {
	go func() {
		for {
			select {
			case <-r.ticker.C:
				log.Println("tick")
			}
		}
	}()
}
