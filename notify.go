package main

import (
    "fmt"
    "syscall"
)


func main() {
    println("start ...")
    fd, err := syscall.InotifyInit()
    if err != nil { panic(err) }
    defer func() {
        err := syscall.Close(fd)
        if err != nil { fmt.Printf("Close = %v", err) }
    }()
    desc, err := syscall.InotifyAddWatch(fd,
                 ".",
                 //"README.md",
                 syscall.IN_MODIFY)
                 //syscall.IN_ALL_EVENTS)
                 //syscall.IN_ACCESS | syscall.IN_DELETE_SELF)
                 //syscall.IN_MODIFY | syscall.IN_ACCESS | syscall.IN_DELETE)
    if err != nil { panic(err) }
    defer func () {
        s, err := syscall.InotifyRmWatch(fd, uint32(desc))
        if err != nil {
            fmt.Print("%v\n", err)
        } else {
            fmt.Print("success=%v\n", s)
        }
    }()
    buf := make([]byte, 100)
    for {
        fmt.Printf("wait to read ..\n")
        n, err := syscall.Read(fd, buf)
        if err != nil { panic(err) }
        fmt.Printf("n=%d, buf[:n]=%v\n", n, buf[:n])
        if n > 16 {
            fmt.Printf("%s\n", string(buf[16:n]))
        }
    }
}
