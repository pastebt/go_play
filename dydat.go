package main

import (
//    "sync"
    "sync/atomic"
)

//https://godoc.org/golang.org/x/exp/inotify


var dat atomic.Value
var dch chan byte = make(chan byte, 1)


func get_data() {
        m := make(map[string]string)
        m["abcd"] = "1234"
        dat.Store(m)
}

func update() {
    for range dch {
        get_data()
    }
}


func main() {
    get_data()
    go update()
    m := dat.Load().(map[string]string)
    println(m["abcd"])
}
