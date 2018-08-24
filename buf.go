package main

import (
    "fmt"
)

func main() {
    buf := make([]byte, 10)
    buf[0] = 'a'
    s := string(buf[:1])
    fmt.Printf("s = [%s]\n", s)
    buf[0] = 'b'
    fmt.Printf("s = [%s]\n", s)
    b2 := []byte(s)
    b2[0] = 'c'
    fmt.Printf("s = [%s]\n", s)
    fmt.Printf("b2[0] = [%c]\n", b2[0])
}
