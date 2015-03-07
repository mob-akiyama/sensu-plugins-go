package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
    url := "http://www.google.co.jp"
  
    resp, _ := http.Get(url)
    defer resp.Body.Close()
  
    byteArray, _ := ioutil.ReadAll(resp.Body)
    fmt.Print(string(byteArray[0:len(byteArray)-1]))
}
