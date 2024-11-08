package parser

import (
	"bantam/expression"
	"bantam/lexer"
	"bantam/token"
	"bantam/tokentype"
	"fmt"
)

/*
One of the two interfaces used by the Pratt parser. A PrefixParselet is
associated with a token that appears at the beginning of an expression. Its
parse() method will be called with the consumed leading token, and the
parselet is responsible for parsing anything that comes after that token.
This interface is also used for single-token expressions like variables, in
which case parse() simply doesn't consume any more tokens.
@author rnystrom
*/
type PrefixParselet interface {
	Parse(*Parser, token.Token) (expression.Expression, error)
}

/*
One of the two parselet interfaces used by the Pratt parser. An
InfixParselet is associated with a token that appears in the middle of the
expression it parses. Its parse() method will be called after the left-hand
side has been parsed, and it in turn is responsible for parsing everything
that comes after the token. This is also used for postfix expressions, in
which case it simply doesn't consume any more tokens in its parse() call.
*/
type InfixParselet interface {
	Parse(*Parser, expression.Expression, token.Token) (expression.Expression, error)
	GetPrecedence() int
}

type Parser struct {
	tokens          *lexer.Lexer // replaced Iterator<Token> of original java code
	read            []token.Token
	prefixParselets map[tokentype.TokenType]PrefixParselet
	infixParselets  map[tokentype.TokenType]InfixParselet
}

/*
Splitted original overloaded java method into two.
You should use these two separately for each Parselets.
*/
func (p *Parser) RegisterPrefix(t tokentype.TokenType, parselet PrefixParselet) {
	p.prefixParselets[t] = parselet
}

func (p *Parser) RegisterInfix(t tokentype.TokenType, parselet InfixParselet) {
	p.infixParselets[t] = parselet
}

func (p *Parser) ParseExpression(precedence int) (expression.Expression, error) {
	token := p.Consume()

	prefix, ok := p.prefixParselets[token.Type]
	if !ok {
		return nil, fmt.Errorf("could not parse \"%s\"", token.Text)
	}

	left, err := prefix.Parse(p, token)
	if err != nil {
		return nil, err
	}

	for precedence < p.getPrecedence() {
		token = p.Consume()
		infix, ok := p.infixParselets[token.Type]
		if !ok {
			return nil, fmt.Errorf("could not find infix parselet for \"%s\"", token.Text)
		}
		var err error
		left, err = infix.Parse(p, left, token)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *Parser) Match(expected tokentype.TokenType) bool {
	if !p.expect(expected) {
		return false
	}
	p.Consume()
	return true
}

func (p *Parser) expect(expected tokentype.TokenType) bool {
	token := p.lookAhead(0)
	return token.Type == expected
}

/*
Splitted original overloaded java method into two.
If you want to do `public Token consume(TokenType expected)`, you should use `ExpectConsume`
*/
func (p *Parser) Consume() token.Token {
	// Make sure we've read the token.
	p.lookAhead(0)

	token := p.read[0]
	p.read = p.read[1:]
	return token
}

func (p *Parser) ExpectConsume(expected tokentype.TokenType) (token.Token, error) {
	t := p.lookAhead(0)
	if t.Type != expected {
		dummyToken := token.Token{}
		return dummyToken, fmt.Errorf("expected token %s and found %s", expected.Punctuator(), t.Type.Punctuator())
	}
	return p.Consume(), nil
}

func (p *Parser) lookAhead(dist int) token.Token {
	// Read in as many as needed
	for dist >= len(p.read) {
		p.read = append(p.read, p.tokens.Next())
	}

	// Get the queued token.
	return p.read[dist]
}

func (p *Parser) getPrecedence() int {
	if parser, ok := p.infixParselets[p.lookAhead(0).Type]; ok {
		return parser.GetPrecedence()
	}
	return 0
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		tokens:          l,
		read:            make([]token.Token, 0),
		prefixParselets: make(map[tokentype.TokenType]PrefixParselet),
		infixParselets:  make(map[tokentype.TokenType]InfixParselet),
	}
}
