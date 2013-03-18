package btcreg

import (
  "fmt"
)

func Main() {
  err := LoadDatabase()
  if err != nil {
    fmt.Println(err)
    return
  }
  a := Address{ "test@example.com", "something bitcoiny" }
  err = InsertAddress(a)
  if err != nil {
    fmt.Println(err)
  }
  RunRouter()
}
