package main

import (
  "unicode"
  "fmt"
)

type TokenKind uint8

const (
  TK_IDENT = iota
  TK_START
  TK_END
  TK_EOF
)

type Token struct {
  Kind TokenKind
  Val string
  Next *Token
}

func NewToken(cur *Token, kind TokenKind, v string) *Token {
  tk := &Token{
    Kind: kind,
    Val: v,
  }
  cur.Next = tk
  return tk
}

type Tokenizer struct {
  s string
  p int
}

func NewTokenizer(s string) *Tokenizer {
  return &Tokenizer{
    s: s,
    p: 0,
  }
}

func (t *Tokenizer) Tokenize() *Token {
  head := &Token{}
  cur := head

  for {
    if len(t.s) == t.p {
      break
    }

    c := t.s[t.p]
    sc := string(c)

    if unicode.IsSpace([]rune(sc)[0]) {
      t.p++
      continue
    }

    fmt.Printf("tk %s \n", sc)

    switch sc {
    case "@":
      t.p++
      val := t.readIdent()
      if val == "startuml" {
        cur = NewToken(cur, TK_START, "@startuml")
      } else if val == "enduml" {
        cur = NewToken(cur, TK_END, "@enduml")
      }
    }
    t.p++
  }

  NewToken(cur, TK_EOF, "EOF")
  return head.Next
}

func (t *Tokenizer) readIdent() string {
  v := ""
  for isLetter(t.s[t.p]) {
    v += string(t.s[t.p])
    t.p++
  }
  return v
}

func isLetter(s byte) bool {
  return 'a' <= s && 'z' >= s
}
