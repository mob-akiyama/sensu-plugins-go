package main

import "fmt" 
import "time"

func main() {
    t := time.Now()
    fmt.Println(t)
    fmt.Println(t.Add(-600 * time.Second))
    fmt.Println(t)
}
