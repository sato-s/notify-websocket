package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	s := &Server{port: port}
	s.run()
	return s
}

func (s *Server) run() {
	app := iris.New()
	ws := neffos.New(websocket.DefaultGorillaUpgrader, neffos.Namespaces{
		"default": neffos.Events{
			"chat":                    s.onChat,
			neffos.OnNamespaceConnect: s.OnRoomJoined,
			neffos.OnRoomJoin:         s.OnRoomJoined,
			"OnRoomJoined":            s.OnRoomJoined,
		},
	})
	app.Get("/", websocket.Handler(ws))

	app.Listen(fmt.Sprintf(":%s", s.port))
}

func (s *Server) onChat(ns *neffos.NSConn, msg neffos.Message) error {
	// ctx := websocket.GetContext(ns.Conn)
	// log.Println(string(msg.Body))
	// fmt.Printf("%+v\n", msg)
	// fmt.Printf("%+v\n", ctx)
	return nil
}

func (s *Server) OnRoomJoined(ns *neffos.NSConn, msg neffos.Message) error {
	ctx := websocket.GetContext(ns.Conn)
	log.Println("====RoomJoined====")
	fmt.Printf("%+v\n", msg)
	fmt.Printf("%+v\n", ctx)
	return nil
}
