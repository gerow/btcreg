package btcreg

import (
  "fmt"
)

func Main() {
  fmt.Println("connecting to mongodb server...")
  err := LoadDatabase()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("connected to mongodb server")
  //a := Address{ "test@example.com", "something bitcoiny" }
  //err = InsertAddress(a)
  //if err != nil {
  //  fmt.Println(err)
  //}
  fmt.Println("initializing router...")
  RunRouter()
}
