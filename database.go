package btcreg

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
   "fmt"
)

var Database *sql.DB

func LoadDatabase() {
  db, err := sql.Open("sqlite3", "./sqlite3.db")
  if err != nil {
    fmt.Println(err)
    return
  }
  db.Exec("create table foo (id integer not null primary key, name text)")
  defer db.Close()
}
