package main

import (
    "fmt"
)


type A struct {
    i, j int
}


type B struct {
    *A
    s string
}


func main () {
    a := A{i: 1, j: 2}
    b := B{A: &a, s: "b_str"}
    fmt.Printf("b = %#v\n", b)
    fmt.Printf("b.i = %#v, b.s = %#v\n", b.i, b.s)
    c := B{A: &a, s: "c_str"}
    fmt.Printf("c.i = %#v, c.s = %#v\n", c.i, c.s)
    a.i = 100
    fmt.Printf("b.i = %#v, b.s = %#v\n", b.i, b.s)
    fmt.Printf("c.i = %#v, c.s = %#v\n", c.i, c.s)
}
