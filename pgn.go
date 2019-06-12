package pgn

import (
	"io"
	"text/scanner"
)

type Tok int

const eof = rune(-1)
const (
	ILLEGAL Tok = iota
	EOF
	WS

	L_BRACE
	R_BRACE
	L_CURLY
	R_CURLY
	DOT

	QUOTE

	IDENT
)

type token struct {
	tok      Tok
	position scanner.Position
	length   int
	literal  string
}

type PGNScanner struct {
	s scanner.Scanner
}

type ParseError struct {
	message  string
	position scanner.Position
}

func (pe ParseError) Error() string {
	return pe.message
}

func (ps *PGNScanner) Next() token {
	char := ps.s.Peek()

	if isWhitespace(char) {
		return ps.scanWhitespace()
	} else if isLetter(char) {
		return ps.scanIdent()
	} else if '"' == char {
		return ps.scanDoubleQuoted()
	}

	ps.s.Next()

	switch char {
	case eof:
		return token{
			tok:      EOF,
			position: ps.s.Pos(),
			length:   1,
		}
	case '[':
		return token{
			tok:      L_BRACE,
			position: ps.s.Pos(),
			length:   1,
		}
	case ']':
		return token{
			tok:      R_BRACE,
			position: ps.s.Pos(),
			length:   1,
		}
	}

	return token{tok: ILLEGAL, position: ps.s.Pos(), length: 1}
}

func (ps *PGNScanner) Init(r io.Reader) {
	ps.s.Init(r)
}
