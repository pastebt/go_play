package main

import (
    "testing"
)


/*
go test -bench . atomic_test.go atomic.go
goos: linux
goarch: amd64
BenchmarkCh-4    	      20	  86818467 ns/op
BenchmarkInt-4   	     100	  21538424 ns/op
*/


func BenchmarkCh(b *testing.B) {
    for i := 0; i < b.N; i++ {
        byChan()
    }
}


func BenchmarkInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        byInt()
    }
}
