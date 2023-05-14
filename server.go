package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

type Server struct {
	port  string
	ws    *neffos.Server
	rooms *Rooms
}

func NewServer(port string) *Server {
	s := &Server{port: port, ws: nil, rooms: nil}

	s.run()
	return s
}

func (s *Server) run() {
	app := iris.New()
	ws := neffos.New(websocket.DefaultGorillaUpgrader, neffos.Namespaces{
		"default": neffos.Events{
			"chat":              s.Debug,
			neffos.OnRoomJoin:   s.OnRoomJoin,
			"startRoomSession":  s.OnStartRoomSession,
			neffos.OnRoomJoined: s.Debug,
		},
	})
	s.ws = ws
	s.rooms = NewRooms(ws)

	app.Get("/", websocket.Handler(s.ws))

	app.Listen(fmt.Sprintf(":%s", s.port))
}

func (s *Server) OnStartRoomSession(ns *neffos.NSConn, msg neffos.Message) error {
	log.Printf("%s", string(msg.Body))

	log.Printf("Session Start at Room `%s Time: %d \n", msg.Room, 10)
	s.rooms.Get(msg.Room).Start()
	return nil
}

func (s *Server) OnRoomJoin(ns *neffos.NSConn, msg neffos.Message) error {
	log.Printf("Room `%s` created\n", msg.Room)
	s.rooms.Get(msg.Room)
	return nil
}

func (s *Server) Debug(ns *neffos.NSConn, msg neffos.Message) error {
	ctx := websocket.GetContext(ns.Conn)
	log.Println("====RoomJoined====")
	fmt.Printf("%+v\n", msg)
	fmt.Printf("%+v\n", ctx)
	return nil
}
