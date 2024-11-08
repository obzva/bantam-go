package lexer

import (
	"bantam/token"
	"bantam/tokentype"
	"unicode"
)

type Lexer struct {
	punctuators map[string]tokentype.TokenType
	index       int
	text        string
}

func (l *Lexer) HasNext() bool { return true }

func (l *Lexer) Next() token.Token {
	for l.index < len(l.text) {
		c := (l.text[l.index])
		l.index++

		if t, ok := l.punctuators[string(c)]; ok {
			// Handle punctuation
			return token.Token{
				Type: t,
				Text: string(c),
			}
		}
		if unicode.IsLetter(rune(c)) {
			// Handle names
			start := l.index - 1
			for l.index < len(l.text) {
				if !unicode.IsLetter(rune(l.text[l.index])) {
					break
				}
				l.index++
			}

			name := l.text[start:l.index]
			return token.Token{
				Type: tokentype.NAME,
				Text: name,
			}
		}
		// Ignore all other characters (whitespace, etc.)
	}

	// Once we've reached the end of the string, just return EOF tokens. We'll
	// just keeping returning them as many times as we're asked so that the
	// parser's lookahead doesn't have to worry about running out of tokens.
	return token.Token{
		Type: tokentype.EOF,
		Text: "",
	}
}

// I skipped the `remove` method from the original java bantam code
// Because Go doesn't support exceptions

func New(t string) *Lexer {
	l := &Lexer{}
	l.text = t
	l.punctuators = make(map[string]tokentype.TokenType)

	// Register all of the TokenTypes that are explicit punctuators.
	for _, tt := range tokentype.TokenTypes {
		p := tt.Punctuator()
		if p != "" {
			l.punctuators[p] = tt
		}
	}

	return l
}
