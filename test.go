package main

import "fmt"

func main() {
    var slice = [...]string{"Penn","Teller"}
    for index := range slice {
        fmt.Print(slice[index])
    }
}
