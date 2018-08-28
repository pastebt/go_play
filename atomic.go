package main

import (
    "sync"
    "sync/atomic"
)


func byChan() {
    var wg sync.WaitGroup
    ch := make(chan int, 1000)
    go func() {
        i := 0
        for {
            i++
            ch <- i
        }
    }()

    wg.Add(4)
    for i := 0; i < 4; i++ {
        go func() {
            for { if c := <-ch; c > 1000000 { break } }
            wg.Done()
        }()
    }
    wg.Wait()
}


func byInt() {
    var wg sync.WaitGroup
    var I int32
    wg.Add(4)
    for i := 0; i < 4; i++ {
        go func() {
            for { if c := atomic.AddInt32(&I, 1); c > 1000000 { break } }
            wg.Done()
        }()
    }
    wg.Wait()
}


func main() {
    byChan()
    byInt()
}
