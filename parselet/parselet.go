package parselet

import (
	"bantam/expression"
	"bantam/parser"
	"bantam/precedence"
	"bantam/token"
	"bantam/tokentype"
	"fmt"
)

// ===== Prefix Parselets =====
/*
Parses parentheses used to group an expression, like "a * (b + c)".
*/
type GroupParselet struct{}

func (gp *GroupParselet) Parse(p *parser.Parser, token token.Token) (expression.Expression, error) {
	exp, err := p.ParseExpression(0)
	if err != nil {
		return nil, err
	}

	_, err = p.ExpectConsume(tokentype.RIGHT_PAREN)
	if err != nil {
		return nil, err
	}

	return exp, nil
}

/*
Simple parselet for a named variable like "abc".
*/
type NameParselet struct{}

func (np *NameParselet) Parse(p *parser.Parser, token token.Token) (expression.Expression, error) {
	nameExp := &expression.NameExpression{
		Name: token.Text,
	}
	return nameExp, nil
}

/*
Generic prefix parselet for an unary arithmetic operator. Parses prefix
unary "-", "+", "~", and "!" expressions.
*/
type PrefixOperatorParselet struct {
	Precedence int
}

func (pp *PrefixOperatorParselet) Parse(p *parser.Parser, token token.Token) (expression.Expression, error) {
	right, err := p.ParseExpression(pp.Precedence)
	if err != nil {
		return nil, err
	}

	prefixExp := &expression.PrefixExpression{
		Operator: token.Type,
		Right:    right,
	}
	return prefixExp, nil
}

func (pp *PrefixOperatorParselet) GetPrecedence() int { return pp.Precedence }

// ===== Infix Parselets =====
/*
Parses assignment expressions like "a = b". The left side of an assignment
expression must be a simple name like "a", and expressions are
right-associative. (In other words, "a = b = c" is parsed as "a = (b = c)").
*/
type AssignParselet struct{}

func (ap *AssignParselet) Parse(p *parser.Parser, left expression.Expression, token token.Token) (expression.Expression, error) {
	right, err := p.ParseExpression(int(precedence.ASSIGNMENT) - 1)
	if err != nil {
		return nil, err
	}

	nameExp, ok := left.(*expression.NameExpression)
	if !ok {
		return nil, fmt.Errorf("the left-hand side of an assignment must be a name")
	}

	assignExp := &expression.AssignExpression{
		Name:  nameExp.Name,
		Right: right,
	}
	return assignExp, nil
}

func (ap *AssignParselet) GetPrecedence() int { return int(precedence.ASSIGNMENT) }

/*
Generic infix parselet for a binary arithmetic operator. The only
difference when parsing, "+", "-", "*", "/", and "^" is precedence and
associativity, so we can use a single parselet class for all of those.
*/
type BinaryOperatorParselet struct {
	Precedence int
	IsRight    bool
}

func (bp *BinaryOperatorParselet) Parse(p *parser.Parser, left expression.Expression, token token.Token) (expression.Expression, error) {
	// To handle right-associative operators like "^", we allow a slightly
	// lower precedence when parsing the right-hand side. This will let a
	// parselet with the same precedence appear on the right, which will then
	// take *this* parselet's result as its left-hand argument.
	decrease := 0
	if bp.IsRight {
		decrease = 1
	}
	right, err := p.ParseExpression(bp.Precedence - decrease)
	if err != nil {
		return nil, err
	}

	operatorExp := &expression.OperatorExpression{
		Left:     left,
		Operator: token.Type,
		Right:    right,
	}
	return operatorExp, nil
}

func (bp *BinaryOperatorParselet) GetPrecedence() int { return bp.Precedence }

/*
Parselet to parse a function call like "a(b, c, d)".
*/
type CallParselet struct{}

func (cp *CallParselet) Parse(p *parser.Parser, left expression.Expression, token token.Token) (expression.Expression, error) {
	// Parse the comma-separated arguments until we hit, ")".
	args := make([]expression.Expression, 0)

	// There may be no arguments at all.
	if !p.Match(tokentype.RIGHT_PAREN) {
		// Do-while like statements
		arg, err := p.ParseExpression(0)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
		for p.Match(tokentype.COMMA) {
			arg, err := p.ParseExpression(0)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}

		_, err = p.ExpectConsume(tokentype.RIGHT_PAREN)
		if err != nil {
			return nil, err
		}
	}

	callExp := &expression.CallExpression{
		Function: left,
		Args:     args,
	}
	return callExp, nil
}

func (cp *CallParselet) GetPrecedence() int { return int(precedence.CALL) }

/*
Parselet for the condition or "ternary" operator, like "a ? b : c".
*/
type ConditionalParselet struct{}

func (cp *ConditionalParselet) Parse(p *parser.Parser, left expression.Expression, token token.Token) (expression.Expression, error) {
	thenArm, err := p.ParseExpression(0)
	if err != nil {
		return nil, err
	}

	_, err = p.ExpectConsume(tokentype.COLON)
	if err != nil {
		return nil, err
	}

	elseArm, err := p.ParseExpression(int(precedence.CONDITIONAL) - 1)
	if err != nil {
		return nil, err
	}

	conditionalExp := &expression.ConditionalExpression{
		Condition: left,
		ThenArm:   thenArm,
		ElseArm:   elseArm,
	}
	return conditionalExp, nil
}

func (cp *ConditionalParselet) GetPrecedence() int { return int(precedence.CONDITIONAL) }

/*
Generic infix parselet for an unary arithmetic operator. Parses postfix
unary "?" expressions.
*/
type PostfixOperatorParselet struct {
	Precedence int
}

func (pp *PostfixOperatorParselet) Parse(p *parser.Parser, left expression.Expression, token token.Token) (expression.Expression, error) {
	postfixExp := &expression.PostfixExpression{
		Left:     left,
		Operator: token.Type,
	}
	return postfixExp, nil
}

func (pp *PostfixOperatorParselet) GetPrecedence() int { return pp.Precedence }
