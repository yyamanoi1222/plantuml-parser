package main

type  Node interface {}

type Statement interface {
  Node
}

type SequenceUml struct {
  Statements []SequenceStatement
}

type SequenceStatement struct {
}


type Uml struct {
  Statements []Statement
}
