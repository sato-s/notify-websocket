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
	ws   *neffos.Server
}

func NewServer(port string) *Server {
	s := &Server{port: port, ws: nil}

	s.run()
	return s
}

func (s *Server) run() {
	app := iris.New()
	ws := neffos.New(websocket.DefaultGorillaUpgrader, neffos.Namespaces{
		"default": neffos.Events{
			"chat":                    s.onChat,
			neffos.OnNamespaceConnect: s.Debug,
			neffos.OnRoomJoin:         s.Debug,
			neffos.OnRoomJoined:       s.Debug,
		},
	})
	fmt.Printf("%+v\n", ws)
	s.ws = ws
	app.Get("/", websocket.Handler(s.ws))

	app.Listen(fmt.Sprintf(":%s", s.port))
}

func (s *Server) onChat(ns *neffos.NSConn, msg neffos.Message) error {
	// ctx := websocket.GetContext(ns.Conn)
	// log.Println(string(msg.Body))
	// fmt.Printf("%+v\n", msg)
	// fmt.Printf("%+v\n", ctx)
	return nil
}

func (s *Server) OnRoomJoin(ns *neffos.NSConn, msg neffos.Message) error {
	ctx := websocket.GetContext(ns.Conn)
	fmt.Printf("%+v\n", msg)
	fmt.Printf("%+v\n", ctx)

	NewRoom("test", s.ws)
	return nil
}

func (s *Server) Debug(ns *neffos.NSConn, msg neffos.Message) error {
	ctx := websocket.GetContext(ns.Conn)
	log.Println("====RoomJoined====")
	fmt.Printf("%+v\n", msg)
	fmt.Printf("%+v\n", ctx)
	return nil
}
