package main

import (
	"bytes"
	"io"
)

type Token int

const (
	WS Token = itoa
	PLUS
	MINUS
	DIVISION
	MULTIPLICATION
	OBRACKET
	CBRACKET
	DIGIT
	EOF
)

//var tokenMap map[rune]Token = {}

type Parser struct {
	b *bytes.Buffer
}

func (p *Parser) newParser(string in) {
	p.b = NewBufferString(in)
}

func (p *Parser) readToken() (token, string) {
	r, size, err := b.ReadRune()
	if err == io.EOF {
		return (EOF, "")
	}

	switch r {
	case " ": return (WS, r)
	case "+": return (PLUS, r)
	case "/":
	}

}
