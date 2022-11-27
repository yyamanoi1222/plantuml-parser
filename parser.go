package main

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
  p.cur = p.l.NextToken()

  for p.cur.Kind != TK_EOF {
    p.cur = p.l.NextToken()
  }
}
