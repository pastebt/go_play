package main

import (
    "os"
)


func main() {
    fi, err := os.Stat("/..")
    println(fi.Name(), fi.Size(), err)
}
