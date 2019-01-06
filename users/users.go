package users

import (
	"goUsers/helper"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type UsersStruct struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Name string        `bson:"name" json:"name"`
}

func session() *mgo.Session {

	//TODO: Add .ENV for this
	//TODO: Is this best practice here?
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session.Clone()

}

//TODO: Im getting 1ms response times on this from my local, is this good?
func All() []UsersStruct {

	defer helper.TimeTrack(time.Now(), "All")

	session := session()
	defer session.Close()

	users := session.DB("usersDB").C("users")

	result := getAll(users)

	return result

}

func getAll(table *mgo.Collection) []UsersStruct {

	var results []UsersStruct
	err := table.Find(bson.M{}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}

	return results

}
