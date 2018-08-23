package main


func f1() (i int) {
    defer func() { if i == 1 { i++ }}()
    return 1
}


func main() {
    println(f1())
}
