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
