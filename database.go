package btcreg

import (
  "labix.org/v2/mgo"
)

var MongoSession *mgo.Session
var DB *mgo.Database

func LoadDatabase() (error) {
	s, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	MongoSession = s
	DB = MongoSession.DB("test")
	err = ensureIndices()
	if err != nil {
		return err
	}
	return nil
}

func ensureIndices() (error) {
	addressIndex := mgo.Index {
		Key: []string{"email"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
	err := DB.C("addresses").EnsureIndex(addressIndex)
	return err
}