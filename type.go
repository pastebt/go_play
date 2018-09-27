package main

import (
    "fmt"
    "encoding/json"
)


func main() {
    var dat interface{}
    err := json.Unmarshal([]byte("[[1, 2, 3], [4, 5, 6]]"), &dat)
    fmt.Printf("%v, %v\n", err, dat)
    switch dat.(type) {
    case []interface{}:
        println("[]interface\n")
        for _, r := range dat.([]interface{}) {
            fmt.Printf("%v\n", r)
        }
    }
}
