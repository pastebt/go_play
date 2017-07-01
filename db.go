package main

import (
    "os"
    "fmt"
    "reflect"
    "database/sql"

 _  "github.com/go-sql-driver/mysql"
)


type Dat struct {
    uid  int64
    ukey string
}


func main() {
    db, err := sql.Open("mysql",
                        //"url:1234@tcp(" + host + ":3306)/url" +
                        os.Args[1] +
                        "?allowOldPasswords=1&parseTime=1&autocommit=true")
    if err != nil { panic(err) }

    uid := 0
    ukey := ""
    q := "select id, ukey from shorturl where ukey > '' and id > ? limit 1"
    if err = db.QueryRow(q, uid).Scan(&uid, &ukey); err != nil { panic(err) }
    println(uid, ukey)

    //uid, ukey = 0, ""
    ds := []interface{}{&uid, &ukey}

    if err = db.QueryRow(q, uid).Scan(ds ...); err != nil { panic(err) }
    println(uid, ukey)

    //ts := []interface{}{reflect.ValueOf(uid).Addr(), &ukey}
    //ts := []interface{}{&uid, reflect.ValueOf(ukey).Addr()}
    dt := new(Dat)
    ts := []interface{}{&dt.uid, &dt.ukey}
    if err = db.QueryRow(q, uid).Scan(ts ...); err != nil { panic(err) }
    fmt.Printf("%v, %v, %v\n", dt, uid, ukey)

    //ts = []interface{}{reflect.ValueOf(dt.uid).Addr(), &dt.ukey}
    ts = []interface{}{reflect.ValueOf(*dt).Field(0), &dt.ukey}
    if err = db.QueryRow(q, dt.uid).Scan(ts ...); err != nil { panic(err) }
    //fmt.Printf("%v, %v, %v\n", dt, uid, ukey)
    fmt.Printf("%v\n", reflect.Kind(uid))
}
