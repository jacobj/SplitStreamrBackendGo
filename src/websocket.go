package main

import (
	"net/http"
)

type WebSocketHandler struct {
	dbs DBStore
}

func (ws WebSocketHandler) serveWs(rw http.ResponseWriter, rq *http.Request) {
	//
}
