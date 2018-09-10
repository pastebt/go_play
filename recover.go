package main

import (
    "fmt"
    "runtime/debug"
)


func rec() {
    if e := recover(); e != nil {
        fmt.Printf("rec.e=%s\n%s\n", e, debug.Stack())
    }
}


func test() (err error) {
    defer func() {
        err = fmt.Errorf("err=%v, recover=%v", err, recover())
    }()
    err = fmt.Errorf("this is a error")
    panic("test")
}


func main() {
    err := test()
    fmt.Printf("test return %v\n", err)
}

