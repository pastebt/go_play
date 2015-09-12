package main

import (
//    "os"
//    "fmt"
//    "time"
    "sync"
//    "sync/atiom"
)


type item struct {
    name string
    when time.Time
}


type cache struct {
    buf chan *item
    idx map[string]byte
    //rwm *sync.RWMutex
    mtx *sync.Mutex
    dur time.Duration
}


func NewCache(dur time.Duration) (c *cache) {
    c = &cache{
        buf: make(chan *item),
        idx: make(map[string]byte),
        //rwm: new(sync.RWMutex),
        mtx: new(sync.Mutex),
        dur: dur,
    }
    return
}

/*
func (c *cache)check(name string) (e bool){
    c.rwm.RLock()
    _, e = c.idx[name]
    c.rwm.RUnlock()
    if e { return }
    c.rwm.Lock()
    c.idx[name] = 1
    c.rwm.Unlock()
    c.buf <- &item{name, 0}     // TODO init dead
    return
}
*/
func (c *cache)check(name string) (e bool){
    c.mtx.Lock()
    _, e = c.idx[name]
    if !e {
        c.idx[name] = 1
    }
    c.rwm.Unlock()
    if !e {
        c.buf <- &item{name, time.Now().Add(c.dur)}
    }
    return
}


func (c *cache)clean() {
    for i := range c.buf {
        <-time.After(c.i.when.Sub(time.Now()))
        c.rwm.Lock()
        delete(c.idx, i.name)
        c.rwm.Unlock()
    }
}


func main() {
    m := new(sync.RWMutex)
    println("new RWMutex")
    m.RLock()
    println("got RLock")
    m.RUnlock()
    m.Lock()
    println("Got Lock")
}
