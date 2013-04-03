package btcreg

import (
"labix.org/v2/mgo/bson"
)

type Address struct {
  Email string
  Address string
}

type AddRequest struct {
  UUID string
  Email string
}

type DeleteRequest struct {
  UUID string
  Email string
}

func LoadAddressByEmail(email string) (Address, error) {
  c := DB.C("addresses")
  result := Address{}
  err := c.Find(bson.M{"email": email}).One(&result)
  if err != nil {
    return result, err
  }
  return result, nil
}

func InsertAddress(address Address) (error) {
  c := DB.C("addresses")
  err := c.Insert(&address)
  if err != nil {
    return err
  }
  return nil
}

//func LoadAddRequestByUUID(uuid string) (AddRequest, error) {
//}

//func LoadDeleteRequestByUUID(uuid string) (DeleteRequest, error) {
//}
