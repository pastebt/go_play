package main
// ex: set tabstop=4 expandtab:
// remember set modeline in vimrc

import (
    "os"
    "fmt"
    "runtime"
    "runtime/debug"
    "strconv"
    "net/http"
    _ "net/http/pprof"
)


type DAT struct {
    dat [1000 * 1000]byte
}


var mm = make(map[int]*DAT)


func hStart(w http.ResponseWriter, req *http.Request) {
    var i int
    s := req.FormValue("cnt")
    i, _ = strconv.Atoi(s)
    for ; i > 0; i-- {
        mm[i] = new(DAT)
    }
	fmt.Fprintf(w, "start %v!", s)
}


func hReset(w http.ResponseWriter, req *http.Request) {
    mm = make(map[int]*DAT)
    act := req.FormValue("act")
    switch act {
    case "gc":
        runtime.GC()
    case "fm":
        debug.FreeOSMemory()
    }
	fmt.Fprintf(w, "hello, world!\n")
}


func main() {
    go http.ListenAndServe(os.Args[1] + "1", nil)
	http.HandleFunc("/start", hStart)
	http.HandleFunc("/reset", hReset)
	http.ListenAndServe(os.Args[1], nil)
}





