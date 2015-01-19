package main

import (
	"gopkg.in/mgo.v2"
)

func createDatabaseSession() *mgo.Session {
	session, error := mgo.Dial("127.0.0.1")

	if error != nil {
		panic(error)
	}

	return session
}
