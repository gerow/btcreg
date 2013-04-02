package btcreg

import (
  "errors"
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
  var a Address 

  query := "SELECT address FROM addresses WHERE email=$1"
  rows, err := Database.Query(query, email)
  if err != nil {
    return a, err
  }
  defer rows.Close()
  var address string 
  if rows.Next() {
    rows.Scan(&address)
  } else {
    return a, errors.New("no such address")
  }

  a.Email = email
  a.Address = address

  return a, nil
}

func InsertAddress(address Address) (error) {
  command := "INSERT INTO addresses (email, address) VALUES ($1, $2)"
  _, err := Database.Exec(command, address.Email, address.Address)
  if err != nil {
    return err
  }
  return nil
}

//func LoadAddRequestByUUID(uuid string) (AddRequest, error) {
//}

//func LoadDeleteRequestByUUID(uuid string) (DeleteRequest, error) {
//}
