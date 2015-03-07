package main

import "fmt"
import "net/http"
import "io/ioutil"

func getInstanceID() string {
    url := "http://169.254.169.254/latest/meta-data/instance-id"
  
    resp, _ := http.Get(url)
    defer resp.Body.Close()
  
    byteArray, _ := ioutil.ReadAll(resp.Body)
    return string(byteArray)
}

func getRegion() string {
    url := "http://169.254.169.254/latest/meta-data/placement/availability-zone"
  
    resp, _ := http.Get(url)
    defer resp.Body.Close()
  
    byteArray, _ := ioutil.ReadAll(resp.Body)
    return string(byteArray)
}

func main() {
    fmt.Println(getInstanceID())
    fmt.Println(getRegion())
}
