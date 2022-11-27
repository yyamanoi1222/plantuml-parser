package main

type TokenKind uint8

const (
  TK_IDENT = iota

  TK_START // @startuml
  TK_END // @enduml

  TK_MINUS // -

  TK_GT // >
  TK_LT // <

  TK_R_ARY // ->
  TK_L_ARY // <-

  TK_COLON // :

  TK_EOF
)

type Token struct {
  Kind TokenKind
  Val string
}

func NewToken(kind TokenKind, v string) *Token {
  return &Token{
    Kind: kind,
    Val: v,
  }
}
