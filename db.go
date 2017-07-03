package main

import (
    "os"
    "fmt"
    "unsafe"
    "reflect"
    "database/sql"
    "database/sql/driver"

_   "github.com/go-sql-driver/mysql"
)


/*********************test Driver***************************/
type tResult struct {
    liid int64
    rwaf int64
}
func (r *tResult)LastInsertId() (int64, error) { return r.liid, nil }
func (r *tResult)RowsAffected() (int64, error) { return r.rwaf, nil }


type tRows struct {
}
func (r *tRows)Columns() []string { return nil }
func (r *tRows)Close() error { return nil }
func (r *tRows)Next(dest []driver.Value) error { return nil }


type tStmt struct {
    num int
}
func (s *tStmt)Close() error { return nil }
func (s *tStmt)NumInput() int { return s.num }
func (s *tStmt)Exec(args []driver.Value) (*tResult, error) {
    t := tResult{liid: 10, rwaf: 1}
    return &t, nil
}
func (s *tStmt)Query(args []driver.Value) (*tRows, error) {
    return &tRows{}, nil
}


type tTx struct {
}
func (t *tTx)Commit() error { return nil }
func (t *tTx)Rollback() error { return nil }


type tstConn struct {
}
func (c *tstConn)Prepare(query string) (*tStmt, error) {
    t := tStmt{}
    return &t, nil
}
func (c *tstConn)Close() error { return nil }
func (c *tstConn)Begin() (*tTx, error) {
    t := tTx{}
    return &t, nil
}


type tstDrv struct {
}
func (t *tstDrv)Open(name string) (*tstConn, error) {
    tc := tstConn{}
    return &tc, nil
}

/*********************test Driver***************************/



type Dat struct {
    uid  int64
    ukey string
}



func main() {
    d := &Dat{uid: 100, ukey: "key"}
    v := reflect.ValueOf(d)
    t := reflect.TypeOf(*d)
    fmt.Printf("t=%v, v=%#v, p=%v, f=%#v\n",
               t, v, v.Pointer(), t.Field(0).Offset)
    a := v.Pointer() + t.Field(1).Offset

    uid := 0
    ukey := ""
    //t2 := reflect.TypeOf(uid)
    t2 := reflect.TypeOf(ukey)
    v2 := reflect.NewAt(t2, unsafe.Pointer(a))
    fmt.Printf("t2=%v, v2=%#v\n", t2, reflect.Indirect(v2))

    if len(os.Args) < 2 { return }
    db, err := sql.Open("mysql",
                        //"url:1234@tcp(" + host + ":3306)/url" +
                        os.Args[1] +
                        "?allowOldPasswords=1&parseTime=1&autocommit=true")
    if err != nil { panic(err) }


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
    //ts = []interface{}{reflect.ValueOf(*dt).Field(0), &dt.ukey}
    //if err = db.QueryRow(q, dt.uid).Scan(ts ...); err != nil { panic(err) }
    //fmt.Printf("%v, %v, %v\n", dt, uid, ukey)
    //fmt.Printf("%v\n", reflect.Kind(uid))
}
