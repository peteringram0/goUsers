package database

import (
	"os"

	"labix.org/v2/mgo"
)

func Session() *mgo.Session {

	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session.Copy()

}
