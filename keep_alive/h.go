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

func svr() {
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%d, Hello, %v\n", cnt, r.RemoteAddr)
        cnt = cnt + 1
        println(r.RemoteAddr)
    })
    println(http.ListenAndServe(":8080", nil))
}


func clt1(to time.Duration) {
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


func clt2(to time.Duration) {
    var conn net.Conn
    t := &http.Transport{Proxy: http.ProxyFromEnvironment,
                         TLSHandshakeTimeout: 10 * time.Second,
                        }
    t.Dial = func(network, addr string) (c net.Conn, e error) {
            if conn != nil { return conn, nil}
            d := &net.Dialer{Timeout: to, KeepAlive: to}
            c, e = d.Dial(network, addr)
            if e == nil { conn = c }
            return
        }
    clt := http.Client{Transport: t}
    // visit first
    resp, err := clt.Get("http://127.0.0.1:8080/bar")
    println(err)
    b := make([]byte, 2, 2)
    n, err := resp.Body.Read(b)
    //buf, err := ioutil.ReadAll(resp.Body)
    println(err)
    println(n)
    println(string(b))
    //println(buf)
    //println(string(buf))
    //resp.Body.Close()

    // visit again
    resp, err = clt.Get("http://127.0.0.1:8080/bar")
    println(err)
    buf, err := ioutil.ReadAll(resp.Body)
    println(err)
    println(buf)
    println(string(buf))
    resp.Body.Close()

    // Done
    conn.Close()
    /*
    resp, err = clt.Get("http://127.0.0.1:8080/bar")
    println(err)
    buf, err = ioutil.ReadAll(resp.Body)
    println(err)
    println(buf)
    println(string(buf))
    resp.Body.Close()
    */
}


func clt3() {
    //t := &http.Transport{Proxy: http.ProxyFromEnvironment,
    //                     TLSHandshakeTimeout: 10 * time.Second,
    //                    }
    //clt := http.Client{Transport: t}
    clt := http.Client{}
    resp, err := clt.Get("http://127.0.0.1:8080/bar")
    println(err.Error())
    if err != nil { return }
    buf, err := ioutil.ReadAll(resp.Body)
    println(err)
    println(buf)
    println(string(buf))
    resp, err = clt.Get("http://127.0.0.1:8080/bar")
    println(err)
    buf, err = ioutil.ReadAll(resp.Body)
    println(err)
    println(buf)
    println(string(buf))
}


func main() {
    if len(os.Args) < 2 {
        svr()
    } else {
        //clt2(time.Duration(10) * time.Second)
        clt3()
    }
}
