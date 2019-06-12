package pgn

import "bytes"

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isNumber(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isIdentChar(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		ch == '=' ||
		ch == '+' ||
		ch == '#' ||
		ch == '-'
}

func (ps *PGNScanner) scanWhitespace() token {
	length := 0
	pos := ps.s.Pos()

	for {
		var char = ps.s.Peek()
		if !isWhitespace(char) {
			break
		} else {
			length += 1
			ps.s.Next()
		}
	}

	return token{
		tok:      WS,
		position: pos,
		length:   length,
	}
}

func (ps *PGNScanner) scanIdent() token {
	length := 0
	pos := ps.s.Pos()
	var buf bytes.Buffer

	for {
		var char = ps.s.Peek()
		if !isIdentChar(char) {
			break
		} else {
			length += 1
			ps.s.Next()
			_, _ = buf.WriteRune(char)
		}
	}

	return token{
		tok:      IDENT,
		position: pos,
		length:   length,
		literal:  buf.String(),
	}
}

func (ps *PGNScanner) scanNumber() token {
	length := 0
	pos := ps.s.Pos()
	var buf bytes.Buffer

	for {
		var char = ps.s.Peek()
		if !isNumber(char) {
			break
		} else {
			length += 1
			ps.s.Next()
			_, _ = buf.WriteRune(char)
		}
	}

	return token{
		tok:      NUMBER,
		position: pos,
		length:   length,
		literal:  buf.String(),
	}
}

func (ps *PGNScanner) scanDoubleQuoted() token {
	ps.s.Next()
	length := 0
	pos := ps.s.Pos()
	var buf bytes.Buffer

	for {
		var char = ps.s.Peek()
		if char == '\\' {
			length += 1
			ps.s.Next()
			c := ps.s.Next()
			_, _ = buf.WriteRune(c)
		} else if char == '"' {
			ps.s.Next()
			break
		} else {
			length += 1
			ps.s.Next()
			_, _ = buf.WriteRune(char)
		}
	}

	return token{
		tok:      QUOTE,
		position: pos,
		length:   length,
		literal:  buf.String(),
	}
}

func (ps *PGNScanner) scanComment() token {
	ps.s.Next()
	length := 0
	pos := ps.s.Pos()
	var buf bytes.Buffer

	for {
		var char = ps.s.Peek()

		if char == '}' {
			ps.s.Next()
			break
		} else {
			length += 1
			ps.s.Next()
			_, _ = buf.WriteRune(char)
		}
	}

	return token{
		tok:      COMMENT,
		position: pos,
		length:   length,
		literal:  buf.String(),
	}
}
