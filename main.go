package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func main() {
	app := iris.New()
	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())

			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	})

	// ws.OnConnect = func(c *websocket.Conn) error {
	// 	log.Printf("[%s] Connected to server!", c.ID())
	// 	return nil
	// }
	//
	// ws.OnDisconnect = func(c *websocket.Conn) {
	// 	log.Printf("[%s] Disconnected from server", c.ID())
	// }

	app.Get("/my_endpoint", websocket.Handler(ws))

	app.Listen(":8080")
}
