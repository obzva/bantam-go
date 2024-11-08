package expression

import (
	"bantam/tokentype"
	"strings"
)

/*
Interface for all expression AST node classes.
*/
type Expression interface {
	// Pretty-print the expression to a string
	Print(*strings.Builder)
}

/*
An assignment expression like "a = b".
*/
type AssignExpression struct {
	Name  string
	Right Expression
}

func (ae *AssignExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	sb.WriteString(ae.Name)
	sb.WriteString(" = ")
	ae.Right.Print(sb)
	sb.WriteString(")")
}

/*
A function call like "a(b, c, d)".
*/
type CallExpression struct {
	Function Expression
	Args     []Expression
}

func (ce *CallExpression) Print(sb *strings.Builder) {
	ce.Function.Print(sb)
	sb.WriteString("(")
	for i, arg := range ce.Args {
		arg.Print(sb)
		if i < (len(ce.Args) - 1) {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
}

/*
A ternary conditional expression like "a ? b : c".
*/
type ConditionalExpression struct {
	Condition Expression
	ThenArm   Expression
	ElseArm   Expression
}

func (ce *ConditionalExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	ce.Condition.Print(sb)
	sb.WriteString(" ? ")
	ce.ThenArm.Print(sb)
	sb.WriteString(" : ")
	ce.ElseArm.Print(sb)
	sb.WriteString(")")
}

/*
A simple variable name expression like "abc".
*/
type NameExpression struct {
	Name string
}

func (ne *NameExpression) Print(sb *strings.Builder) {
	sb.WriteString(ne.Name)
}

/*
A binary arithmetic expression like "a + b" or "c ^ d".
*/
type OperatorExpression struct {
	Left     Expression
	Operator tokentype.TokenType
	Right    Expression
}

func (oe *OperatorExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	oe.Left.Print(sb)
	sb.WriteString(" ")
	sb.WriteString(oe.Operator.Punctuator())
	sb.WriteString(" ")
	oe.Right.Print(sb)
	sb.WriteString(")")
}

/*
A postfix unary arithmetic expression like "a!".
*/
type PostfixExpression struct {
	Left     Expression
	Operator tokentype.TokenType
}

func (pe *PostfixExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	pe.Left.Print(sb)
	sb.WriteString(pe.Operator.Punctuator())
	sb.WriteString(")")
}

/*
A prefix unary arithmetic expression like "!a" or "-b".
*/
type PrefixExpression struct {
	Operator tokentype.TokenType
	Right    Expression
}

func (pe *PrefixExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	sb.WriteString(pe.Operator.Punctuator())
	pe.Right.Print(sb)
	sb.WriteString(")")
}
