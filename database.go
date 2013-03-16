package btcreg

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
   "fmt"
)

var Database *sql.DB
var DBVersion int = 1

func LoadDatabase() {
  db, err := sql.Open("sqlite3", "./btcreg.sqlite3")
  if err != nil {
    fmt.Println(err)
    return
  }
  Database = db
}
