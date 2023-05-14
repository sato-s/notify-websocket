package main

import (
	"sync"

	"github.com/kataras/neffos"
)

type Rooms struct {
	list sync.Map
	ws   *neffos.Server
}

func NewRooms(ws *neffos.Server) *Rooms {
	rooms := &Rooms{list: sync.Map{}, ws: ws}

	return rooms
}

func (r *Rooms) Get(name string) *Room {
	if room, ok := r.list.Load(name); ok {
		return room.(*Room)
	} else {
		newRoom := NewRoom(name, r.ws)
		r.list.Store(name, newRoom)
		return newRoom
	}
}
