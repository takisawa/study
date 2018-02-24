package main

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  fmt.Println("test")

  db, err := sql.Open("mysql", "root@localhost:3306/test_db")
	if err != nil {
    fmt.Println(err)
    return
	}

  fmt.Println(db)
}
