package main

import "fmt"
import "net/http"
import "io/ioutil"

func main() {
  url := "http://169.254.169.254/latest/meta-data/instance-id"

  resp, _ := http.Get(url)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray)) // htmlをstringで取得
}
