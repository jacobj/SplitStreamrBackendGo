package main

import (
	"gopkg.in/mgo.v2"
)

type DBStore struct {
	session *mgo.Session
	name    string
}

func (dbs *DBStore) songs() *mgo.Collection {
	return dbs.session.DB(dbs.name).C("songs")
}
