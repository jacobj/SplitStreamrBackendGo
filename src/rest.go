package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Song struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	Name           string        `bson"name" json:"name"`
	Artist         string        `bson"artist" json:"artist"`
	Length         int64         `bson"length" json:"length"`
	NumberOfChunks int64         `bson"numberOfChunks" json:"numberOfChunks"`
	FileType       string        `bson"fileType" json:"fileType"`
	FileSize       int64         `bson"fileSize" json:"fileSize"`
	Path           string        `bson"path" json:"path"`
}

type RestHandler struct {
	dbs DBStore
}

func (rh RestHandler) handleSong(rw http.ResponseWriter, rq *http.Request) {
	id := rq.URL.String()[len("/songs/"):]
	if id == "" {
		// No songId, i.e. bad request.
		http.Error(rw, "Bad request", 400)
	}
	switch rq.Method {
	case "GET":
		result := Song{}
		rh.dbs.songs().FindId(bson.ObjectIdHex(id)).One(&result)

		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		song, _ := json.Marshal(result)

		rw.Write(song)
		return
	default:
		http.Error(rw, "Method not allowed", 405)
	}
}

func (rh RestHandler) handleSongs(rw http.ResponseWriter, rq *http.Request) {
	switch rq.Method {
	case "GET":
		result := []Song{}
		rh.dbs.songs().Find(bson.M{}).All(&result)

		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		songs, _ := json.Marshal(result)

		rw.Write(songs)
		return
	default:
		http.Error(rw, "Method not allowed", 405)
	}
}
