package users

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type usersStruct struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}

func session() *mgo.Session {

	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session

}

func SayHi() {

	session := session()
	defer session.Close()

	users := session.DB("usersDB").C("users")

	result := getAll(users)
	fmt.Printf("%v", result)

}

func getAll(table *mgo.Collection) []usersStruct {
	var results []usersStruct
	err := table.Find(bson.M{}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	return results
}
