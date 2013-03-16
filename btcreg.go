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
  RunRouter()
}
