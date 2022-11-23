package main

import (
  "fmt"
  "os"
  "io/ioutil"
)

func main() {
  body, _ := ioutil.ReadAll(os.Stdin)
  t := NewTokenizer(string(body))
  token := t.Tokenize()

  for token != nil {
    fmt.Printf("tk: %v \n", token)
    token = token.Next
  }
}
