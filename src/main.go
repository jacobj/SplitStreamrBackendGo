package main

import (
	"flag"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal("mgo.Dial: ", err)
	}

	defer session.Close()

	dbs := DBStore{name: "splitstreamr-test", session: session}

	rh := RestHandler{dbs: dbs}

	wsh := WebSocketHandler{dbs: dbs}

	http.HandleFunc("/songs/", rh.handleSong)
	http.HandleFunc("/songs", rh.handleSongs)
	http.HandleFunc("/ws", wsh.serveWs)

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
