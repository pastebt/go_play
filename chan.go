package main

import (
    "fmt"
)


func main() {

    a := make(chan int, 3)
    b := make(chan byte, 3)
    a <- 1
    a <- 2
    b <- 3
    b <- 4
    close(a)
    close(b)

    for cnt := 0; cnt < 4; {
        select {
        case A, ok := <-a:
            if ok {
                cnt += 1
                fmt.Printf("A = %v, ok=%v\n", A, ok)
            }
        case B, ok := <-b:
            if ok {
                cnt += 1
                fmt.Printf("B = %v, ok=%v\n", B, ok)
            }
        }
    }
}
