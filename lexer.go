package main

import (
  "fmt"
)

type Lexer struct {
  s string
  p int
  ch byte
}

func NewLexer(s string) *Lexer {
  l := &Lexer{
    p: 0,
    s: s,
  }
  l.readChar()
  return l
}

func (l *Lexer) NextToken() *Token {
  l.skipWhitespace()
  var tk *Token

  ch := l.ch
  switch ch {
    case '@':
      val := l.readIdent()
      if val == "startuml" {
        tk = NewToken(TK_START, "@startuml")
      } else if val == "enduml" {
        tk = NewToken(TK_END, "@enduml")
      }
    case '-':
      if l.peekChar() == '>' {
        l.readChar()
        tk = NewToken(TK_R_ARY, "->")
      } else {
        tk = NewToken(TK_MINUS, "-")
      }
    case '>':
      tk = NewToken(TK_GT, ">")
    case '<':
      if l.peekChar() == '-' {
        l.readChar()
        tk = NewToken(TK_L_ARY, "<-")
      } else {
        tk = NewToken(TK_LT, "<")
      }
    case ':':
      tk = NewToken(TK_COLON, ":")
    case 0:
      tk = NewToken(TK_EOF, "EOF")
    default:
      val := l.readIdent()
      tk = NewToken(TK_IDENT, string(ch) + val)
  }
  fmt.Printf("%v\n", tk)
  l.readChar()
  return tk
}

func (l *Lexer) peekChar() byte {
  if len(l.s) <= l.p {
    return 0
  } else {
    return l.s[l.p]
  }
}

func (l *Lexer) readChar() {
  if len(l.s) <= l.p {
    l.ch = 0
  } else {
    l.ch = l.s[l.p]
  }
  l.p++
}

func (l *Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\n' {
    l.readChar()
  }
}

func (l *Lexer) readIdent() string {
  v := ""
  for isLetter(l.s[l.p]) {
    v += string(l.s[l.p])
    l.readChar()
  }
  return v
}

func isLetter(s byte) bool {
  return 'a' <= s && s <= 'z' || 'A' <= s && s <= 'Z'
}
