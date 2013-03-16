package main

import(
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/gerow/btcreg/migrate_db"
    "github.com/gerow/btcreg"
)

var migrations map[int] migratedb.Migration = map[int] migratedb.Migration{
  // First migration
  1 : migratedb.Migration{
    func(tx *sql.Tx) (error) {
      sql := `CREATE TABLE addresses (
          id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
          email VARCHAR NOT NULL,
          address VARCHAR NOT NULL,
          created DATETIME NULL DEFAULT CURRENT_TIMESTAMP
          );`
      fmt.Println(sql)
      _, err := tx.Exec(sql)
      if err != nil {
        return err
      }
      sql = `CREATE TABLE add_requests (
          id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
          email VARCHAR NOT NULL,
          uuid VARCHAR NOT NULL,
          created DATETIME NULL DEFAULT CURRENT_TIMESTAMP
          );`
      _, err = tx.Exec(sql)
      if err != nil {
        return err
      }
      sql = `CREATE TABLE delete_requests (
          id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
          email VARCHAR NOT NULL,
          uuid VARCHAR NOT NULL,
          created DATETIME NULL DEFAULT CURRENT_TIMESTAMP
          );`
      _, err = tx.Exec(sql)
      if err != nil {
        return err 
      }
      return nil
    },
  },
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

    version, err := migratedb.Version(tx)
    if version == -1 {
      fmt.Println("db is uninitialized, creating version table")
      migratedb.CreateVersionTable(tx)
      version, err = migratedb.Version(tx)
    }
    if err != nil {
      migratedb.Rollback(tx)
      fmt.Println(err)
      return
    }
    err = tx.Commit()
    if err != nil {
      fmt.Println(err)
      return
    }
    tx, err = db.Begin()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Printf("target DB version is %d\n", btcreg.DBVersion)
    fmt.Printf("currently on version %d\n", version)
    if version == btcreg.DBVersion {
      fmt.Printf("db is up to date. exiting.\n")
      return
    }
    for version != btcreg.DBVersion {
      version += 1
      fmt.Printf("migrating to version %d\n", version)
      err = migrations[version].Up(tx)
      if err != nil {
        fmt.Println(err)
        migratedb.Rollback(tx)
        fmt.Println("failed to migrate database")
        return
      }
      migratedb.UpdateVersion(tx, version)
    }
    err = tx.Commit()
    if err != nil {
      fmt.Println(err)
      return
    }
}
