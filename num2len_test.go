package main

import (
    "testing"
    "fmt"
    "math"
)


func num2len(n int64) (l int) {
    l = 1
    for ; n > 9; n = n / 10 {
        l += 1
    }
    return
}


func fmt2str(n int64) (l int) {
    s := fmt.Sprintf("%d", n)
    return len(s)
}



func BenchmarkNum2len(bm *testing.B) {
    var n int64 = 12345
    for i := 0; i < bm.N; i++ {
        num2len(n)
    }
}


func BenchmarkLog10(bm *testing.B) {
    var n int64 = 12345
    for i := 0; i < bm.N; i++ {
        math.Log10(float64(n))
    }
}


func BenchmarkFmt2str(bm *testing.B) {
    var n int64 = 12345
    for i := 0; i < bm.N; i++ {
        fmt2str(n)
    }
}


func TestNum(tst *testing.T) {
    ns := []int64{1,      1,
                  0,      1,
                  12,     2,
                  321,    3,
                  5443,   4,
                  98912,  5,
                  123457, 6,
                 }
    for i := 0; i < len(ns); i += 2 {
        if int64(num2len(ns[i])) != ns[i + 1] {
            tst.Errorf("Wrong len(%d) = %d != %d",
                       ns[i], num2len(ns[i]), ns[i + 1])
        }
    }
}
