package btcreg

import (
  "labix.org/v2/mgo"
  "fmt"
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
	fmt.Println("creating indices if they don't already exist...")
	err = ensureIndices()
	if err != nil {
		return err
	}
	fmt.Println("done creating indices")
	return nil
}

func ensureIndices() (error) {
	addressIndex := mgo.Index {
		Key: []string{"email"},
		Unique: true,
		DropDups: true,
		Background: false,
		Sparse: true,
	}
	err := DB.C("addresses").EnsureIndex(addressIndex)
	return err
}