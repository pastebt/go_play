package main

import (
//    "os"
    "fmt"
//    "strings"
    "database/sql"
_   "rsc.io/sqlite"
)


func main() {
    db, err := sql.Open("sqlite3", "/tmp/test.db")
    if err != nil { panic(err) }
    fmt.Printf("db = %#v\n", db)
    /*
    res, err := db.Exec("create table a (uid int);")
    if err != nil {
        if ! strings.Contains(err.Error(), "table a already exists") {
            panic(err)
        }
    }
    fmt.Printf("res = %#v\n", res)
    */
    // total ignore error, will hit again in future sql opt
    db.Exec("create table a (uid int);")
    res, err := db.Exec("insert into a (uid) values (?)", 100)
    if err != nil { panic(err) }
    fmt.Printf("res = %#v\n", res)
}
