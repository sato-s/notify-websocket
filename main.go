package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

func onChat(ns *neffos.NSConn, msg neffos.Message) error {
	// ctx := websocket.GetContext(ns.Conn)
	// log.Println(string(msg.Body))
	// fmt.Printf("%+v\n", msg)
	// fmt.Printf("%+v\n", ctx)
	return nil
}

func OnRoomJoined(ns *neffos.NSConn, msg neffos.Message) error {
	ctx := websocket.GetContext(ns.Conn)
	log.Println("====RoomJoined====")
	fmt.Printf("%+v\n", msg)
	fmt.Printf("%+v\n", ctx)
	return nil
}

func main() {
	app := iris.New()
	ws := neffos.New(websocket.DefaultGorillaUpgrader, neffos.Namespaces{
		"default": neffos.Events{
			"chat":                    onChat,
			neffos.OnNamespaceConnect: OnRoomJoined,
			neffos.OnRoomJoin:         OnRoomJoined,
			"OnRoomJoined":            OnRoomJoined,
		},
	})
	app.Get("/", websocket.Handler(ws))

	app.Listen(":8080")
}
