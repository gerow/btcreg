package migratedb

import(
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Migration struct {
  Up func(db *sql.Tx) (error)
}

func CreateVersionTable(tx *sql.Tx) error {
  sql := "CREATE TABLE db_version (id INTEGER NOT NULL PRIMARY KEY, version INTEGER NOT NULL)"
  _, err := tx.Exec(sql)
  if err != nil {
    return err
  }
  _, err = tx.Exec("INSERT INTO db_version(id, version) values(0, 0)")
  if err != nil {
    return err
  }
  return nil
}

func Version(tx *sql.Tx) (int, error) {
  rows, err := tx.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='db_version'")
  if err != nil {
    return -1, err
  }
  hasTable := rows.Next()
  rows.Close()
  if !hasTable {
    return -1, nil
  }
  rows, err = tx.Query("SELECT version FROM db_version WHERE id=0")
  rows.Next()
  var version int
  rows.Scan(&version)
  return version, nil
}

func UpdateVersion(tx *sql.Tx, version int) (error) {
  sql := "UPDATE db_version SET version=$1"
  _, err := tx.Exec(sql, version)
  if err != nil {
    return err
  }
  return nil
}

func Rollback(tx *sql.Tx) {
  fmt.Println("rolling back changes")
  err := tx.Rollback()
  if err != nil {
    fmt.Println(err)
    return
  }
}
