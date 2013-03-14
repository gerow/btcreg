package main

import(
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/gerow/btcreg"
)

type Migration struct {
  Up func(db *sql.DB) (success bool)
}

var migrations map[int] Migration = map[int] Migration{
  // First migration
  1 : Migration{
    func(db *sql.DB) (success bool) {
      return true
    },
  },
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

func main() {
    driver := btcreg.Conf.SqlDriver()
    db, err := sql.Open(driver, btcreg.Conf.SqlOpen())
    if err != nil {
      fmt.Println(err)
      return
    }
    tx, err := db.Begin()
    if err != nil {
      fmt.Println(err)
      return
    }

    version, err := Version(tx)
    if version == -1 {
      fmt.Println("db is uninitialized, creating version table")
      CreateVersionTable(tx)
      version, err = Version(tx)
    }
    if err != nil {
      Rollback(tx)
      fmt.Println(err)
    }
    fmt.Printf("target DB version is %d\n", btcreg.DBVersion)
    fmt.Printf("currently on version %d\n", version)
    if version == btcreg.DBVersion {
      fmt.Printf("db is up to date. exiting.\n")
    }
    err = tx.Commit()
    if err != nil {
      fmt.Println(err)
      return
    }
}

func Rollback(tx *sql.Tx) {
  fmt.Println("rolling back changes")
  err := tx.Rollback()
  if err != nil {
    fmt.Println(err)
    return
  }
}
