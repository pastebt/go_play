package main

import (
//    "os"
    "fmt"
//    "net"
//    "time"
    "net/http"
//    "io/ioutil"
)


var cnt int = 0

func svr1() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%d, Hello, %v, %v\n", cnt, r.RemoteAddr, r.RequestURI)
        cnt = cnt + 1
        println(r.RemoteAddr)
    })
    println(http.ListenAndServe(":8080", nil))
}


func wk(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "wk: %d, Hello, %v, %v\n", cnt, r.RemoteAddr, r.RequestURI)
    cnt = cnt + 1
    print(r.RemoteAddr)
    println(" " + r.RequestURI)
}


func svr2() {
    s := http.Server{Addr: ":8080", Handler: http.HandlerFunc(wk)}
    s.ListenAndServe()
}


func main() {
    svr2()
}
