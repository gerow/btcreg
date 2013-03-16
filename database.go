package btcreg

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB
var DBVersion int = 1

func LoadDatabase() (error) {
  db, err := sql.Open("sqlite3", "./btcreg.sqlite3")
  if err != nil {
    return err
  }
  Database = db
  return nil
}
