package users

import (
	"goUsers/database"
	"goUsers/helper"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type UsersStruct struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Name string        `bson:"name" json:"name"`
}

func All() []UsersStruct {

	defer helper.TimeTrack(time.Now(), "All")

	session := database.Session()
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
