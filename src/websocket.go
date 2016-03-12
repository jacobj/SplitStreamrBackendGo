package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// Upgrader to upgrade HTTP request in serveWs
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type WebSocketHandler struct {
	dbs         DBStore
	connections map[*connection]bool
	broadcast   chan []byte
	register    chan *connection
	unregister  chan *connection
}

func (wsh *WebSocketHandler) run() {
	for {
		select {
		case c := <-wsh.register:
			wsh.connections[c] = true
		case c := <-wsh.unregister:
			if _, ok := wsh.connections[c]; ok {
				delete(wsh.connections, c)
				close(c.send)
			}
		case m := <-wsh.broadcast:
			for c := range wsh.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(wsh.connections, c)
				}
			}
		}
	}
}

func (wsh *WebSocketHandler) serveWs(rw http.ResponseWriter, rq *http.Request) {
	switch rq.Method {
	case "GET":
		ws, err := upgrader.Upgrade(rw, rq, nil)
		if err != nil {
			log.Println(err)
			return
		}
		c := &connection{send: make(chan []byte, 256), ws: ws}
		wsh.register <- c
	default:
		http.Error(rw, "Method not allowed", 405)
	}
}
