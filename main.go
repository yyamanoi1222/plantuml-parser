package main

import (
  "os"
  "io/ioutil"
)

func main() {
  body, _ := ioutil.ReadAll(os.Stdin)
  lexer := NewLexer(string(body))
  parser := NewParser(lexer)
  parser.Parse()
}
