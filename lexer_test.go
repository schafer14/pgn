package pgn

import (
	"strings"
	"testing"
)

type scannerTest struct {
	phrase string
	tokens []token
	name   string
}

var testers = []scannerTest{
	scannerTest{
		name:   "Tag Happy Case",
		phrase: `[White "Fabiano Caruana"]`,
		tokens: []token{
			token{
				tok: L_BRACE,
			},
			token{
				tok:     IDENT,
				literal: "White",
			},
			token{
				tok: WS,
			},
			token{
				tok:     QUOTE,
				literal: "Fabiano Caruana",
			},
			token{
				tok: R_BRACE,
			},
			token{
				tok: EOF,
			},
		},
	},
	scannerTest{
		name:   `Tag with \\`,
		phrase: `[White "Fabiano \\Caruana"]`,
		tokens: []token{
			token{
				tok: L_BRACE,
			},
			token{
				tok:     IDENT,
				literal: "White",
			},
			token{
				tok: WS,
			},
			token{
				tok:     QUOTE,
				literal: `Fabiano \Caruana`,
			},
			token{
				tok: R_BRACE,
			},
			token{
				tok: EOF,
			},
		},
	},
	scannerTest{
		name:   `Tag with \"`,
		phrase: `[White "Fabiano \"Caruana"]`,
		tokens: []token{
			token{
				tok: L_BRACE,
			},
			token{
				tok:     IDENT,
				literal: "White",
			},
			token{
				tok: WS,
			},
			token{
				tok:     QUOTE,
				literal: `Fabiano "Caruana`,
			},
			token{
				tok: R_BRACE,
			},
			token{
				tok: EOF,
			},
		},
	},
	scannerTest{
		name:   `Happy Moves`,
		phrase: `1. e4 c5 2. Nf6`,
		tokens: []token{
			token{
				tok:     NUMBER,
				literal: "1",
			},
			token{
				tok: DOT,
			},
			token{
				tok: WS,
			},
			token{
				tok:     IDENT,
				literal: `e4`,
			},
			token{
				tok: WS,
			},
			token{
				tok:     IDENT,
				literal: `c5`,
			},
			token{
				tok: WS,
			},
			token{
				tok:     NUMBER,
				literal: "2",
			},
			token{
				tok: DOT,
			},
			token{
				tok: WS,
			},
			token{
				tok:     IDENT,
				literal: "Nf6",
			},
		},
	},
	scannerTest{
		name:   `Comments`,
		phrase: `1. e4 {Incredible}`,
		tokens: []token{
			token{
				tok:     NUMBER,
				literal: "1",
			},
			token{
				tok: DOT,
			},
			token{
				tok: WS,
			},
			token{
				tok:     IDENT,
				literal: `e4`,
			},
			token{
				tok: WS,
			},
			token{
				tok:     COMMENT,
				literal: `Incredible`,
			},
		},
	},
	scannerTest{
		name:   `Alternatives`,
		phrase: `1. e4 (1. d4)`,
		tokens: []token{
			token{
				tok:     NUMBER,
				literal: "1",
			},
			token{
				tok: DOT,
			},
			token{
				tok: WS,
			},
			token{
				tok:     IDENT,
				literal: `e4`,
			},
			token{
				tok: WS,
			},
			token{
				tok: L_PAREN,
			},
			token{
				tok:     NUMBER,
				literal: "1",
			},
			token{
				tok: DOT,
			},
			token{
				tok: WS,
			},
			token{
				tok:     IDENT,
				literal: "d4",
			},
			token{
				tok: R_PAREN,
			},
		},
	},
}

// TODO: test phrases with \\ and \" characters
func TestNext(t *testing.T) {

	for _, tester := range testers {
		var s PGNScanner

		s.Init(strings.NewReader(tester.phrase))

		for i, expected := range tester.tokens {
			got := s.Next()

			if got.tok != expected.tok {
				t.Errorf("Error in test %s token %d: expected tok %v but got %v", tester.name, i, expected.tok, got.tok)
			}

			if got.literal != expected.literal {
				t.Errorf("Error in test %s token %d: expected literal %s but got %s", tester.name, i, expected.literal, got.literal)
			}
		}
	}
}
