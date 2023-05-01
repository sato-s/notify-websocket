package main

import (
	"log"
	"time"
)

type Room struct {
	name   string
	ticker *time.Ticker
}

func NewRoom(name string) *Room {
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
