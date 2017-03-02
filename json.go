package main

import (
    "fmt"
    "encoding/json"
)


type A struct {
    S string
}

type B struct {
    A
    I int
}


func main() {
    b := B{A{"a_str"}, 10}
    m, e := json.Marshal(b)

    fmt.Printf("b=%v, e=%v\n", string(m), e)
    var a A
    e = json.Unmarshal(m, &a)
    fmt.Printf("a=%v, e=%v, S=%v\n", a, e, a.S)

    a.S = "str2"
    n, e := json.Marshal(a)
    fmt.Printf("n=%v, e=%v\n", string(n), e)
    e = json.Unmarshal(n, &b)
    fmt.Printf("b=%v, e=%v, S=%v, I=%v\n", b, e, b.S, b.I)
}
