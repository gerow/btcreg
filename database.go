package btcreg

import (
  "labix.org/v2/mgo"
)

var MongoSession *mgo.Session

func LoadDatabase() (error) {
	s, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	MongoSession = s
	return nil
}
