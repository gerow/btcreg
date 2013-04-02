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
	err := MongoSession.DB("test").C("addresses").EnsureIndex(addressIndex)
	return err
}