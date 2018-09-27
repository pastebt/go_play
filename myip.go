package main

import (
    "fmt"
    "regexp"
    "net/http"
    "io/ioutil"
)


func main() {
    res, err := http.Get("https://www.myip.com/")
    if err != nil { panic(err) }
    defer res.Body.Close()
    dat, err := ioutil.ReadAll(res.Body)
    if err != nil { panic(err) }
    msg := string(dat)
    fmt.Printf("%v\n", msg)
    m, err := regexp.Compile(`\<span id\="ip"\>([0-9.]+)\</span\>`)
    if err != nil { panic(err) }
    mc := m.FindStringSubmatch(msg)
    fmt.Printf("%v\n", mc)
}
