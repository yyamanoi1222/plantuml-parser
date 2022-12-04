package main

import (
  "log"
)

type Parser struct {
  l *Lexer
  cur *Token
}

func NewParser(l *Lexer) *Parser {
  return &Parser{
    l: l,
  }
}

func (p *Parser) Parse() {
  p.nextToken()
  uml := p.parseUml()
  if uml == nil {
    log.Fatal("invalid uml exit")
  }
}

func (p *Parser) nextToken() {
  p.cur = p.l.NextToken()
}

func (p *Parser) parseUml() *Uml {
  uml := &Uml{}
  if p.cur.Kind != TK_START {
    return nil
  }
  b := p.cur

  for {
    stmt := p.parseStatement()
    if stmt == nil {
      break
    }

    if p.cur.Kind == TK_EOF {
      break
    }
    b = p.cur
  }

  if b.Kind != TK_END {
    return nil
  }

  return uml
}

func (p *Parser) parseStatement() *Statement {
  switch p.cur.Kind {
  case TK_IDENT:
    return p.parseUmlDef()
  }
  log.Fatal("invalid statement")
  return nil
}

func (p *Parser) parseUmlDef() *Statement {
  switch p.l.peekChar() {
  case TK_R_ARY:
    return p.parseSequence()
  }
  log.Fatal("invalid uml def")
  return nil
}

func (p *Parser) parseSequence() *Statement {
}
