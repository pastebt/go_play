package main


import (
    "testing"
)


func TestTest(tst *testing.T) {
    s := test()
    if s != "test string" {
        tst.Error(s)
    }
}
