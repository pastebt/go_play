package main


import (
    "os"
    "fmt"
    "net"
    "time"
    "net/http"
    "io/ioutil"
)


var cnt int = 0

func bar(w http.ResponseWriter, r *http.Request) {
    //time.Sleep(10 * time.Second)
    fmt.Fprintf(w, "%d, Hello, %v\n", cnt, r.RemoteAddr)
    cnt = cnt + 1
    println(r.RemoteAddr)
}


func svr() {
    http.HandleFunc("/bar", bar)
    println(http.ListenAndServe(":8080", nil))
}


func clt(to time.Duration) {
    var conn net.Conn
    t := &http.Transport{Proxy: http.ProxyFromEnvironment,
                         TLSHandshakeTimeout: 10 * time.Second,
                        }
    t.Dial = func(network, addr string) (c net.Conn, e error) {
            d := &net.Dialer{Timeout: to, KeepAlive: to}
            c, e = d.Dial(network, addr)
            if e == nil { conn = c }
            return
        }
    clt := http.Client{Transport: t}
    // visit fup, get task
    resp, err := clt.Get("http://127.0.0.1:8080/bar")
    println(err)
    buf, err := ioutil.ReadAll(resp.Body)
    println(err)
    println(buf)
    println(string(buf))
    resp.Body.Close()

}


func main() {
    if len(os.Args) < 2 {
        svr()
    } else {
        clt(time.Duration(5) * time.Second)
    }
}
