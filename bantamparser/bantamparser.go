package bantamparser

import (
	"bantam/lexer"
	"bantam/parselet"
	"bantam/parser"
	"bantam/precedence"
	"bantam/tokentype"
)

/*
Extends the generic Parser class with support for parsing the actual Bantam
grammar.
*/
func New(l *lexer.Lexer) *parser.Parser {
	p := parser.New(l)

	// Register all of the parselets for the grammar.

	// Register the ones that need special parselets.
	p.RegisterPrefix(tokentype.NAME, &parselet.NameParselet{})
	p.RegisterInfix(tokentype.ASSIGN, &parselet.AssignParselet{})
	p.RegisterInfix(tokentype.QUESTION, &parselet.ConditionalParselet{})
	p.RegisterPrefix(tokentype.LEFT_PAREN, &parselet.GroupParselet{})
	p.RegisterInfix(tokentype.LEFT_PAREN, &parselet.CallParselet{})

	// Register the simple operator parselets.
	p.RegisterPrefix(tokentype.PLUS, &parselet.PrefixOperatorParselet{Precedence: int(precedence.PREFIX)})
	p.RegisterPrefix(tokentype.MINUS, &parselet.PrefixOperatorParselet{Precedence: int(precedence.PREFIX)})
	p.RegisterPrefix(tokentype.TILDE, &parselet.PrefixOperatorParselet{Precedence: int(precedence.PREFIX)})
	p.RegisterPrefix(tokentype.BANG, &parselet.PrefixOperatorParselet{Precedence: int(precedence.PREFIX)})

	// For kicks, we'll make "!" both prefix and postfix, kind of like ++.
	p.RegisterInfix(tokentype.BANG, &parselet.PostfixOperatorParselet{Precedence: int(precedence.POSTFIX)})

	p.RegisterInfix(tokentype.PLUS, &parselet.BinaryOperatorParselet{Precedence: int(precedence.SUM), IsRight: false})
	p.RegisterInfix(tokentype.MINUS, &parselet.BinaryOperatorParselet{Precedence: int(precedence.SUM), IsRight: false})
	p.RegisterInfix(tokentype.ASTERISK, &parselet.BinaryOperatorParselet{Precedence: int(precedence.PRODUCT), IsRight: false})
	p.RegisterInfix(tokentype.SLASH, &parselet.BinaryOperatorParselet{Precedence: int(precedence.PRODUCT), IsRight: false})
	p.RegisterInfix(tokentype.CARET, &parselet.BinaryOperatorParselet{Precedence: int(precedence.EXPONENT), IsRight: true})

	return p
}
