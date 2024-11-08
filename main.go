package main

import (
	"bantam/bantamparser"
	"bantam/lexer"
	"fmt"
	"strings"
)

var (
	passed = 0
	failed = 0
)

func main() {
	// Function call
	test("a()", "a()")
	test("a(b)", "a(b)")
	test("a(b, c)", "a(b, c)")
	test("a(b)(c)", "a(b)(c)")
	test("a(b) + c(d)", "(a(b) + c(d))")
	test("a(b ? c : d, e + f)", "a((b ? c : d), (e + f))")

	// Unary precedence
	test("~!-+a", "(~(!(-(+a))))")
	test("a!!!", "(((a!)!)!)")

	// Unary and binary precedence
	test("-a * b", "((-a) * b)")
	test("!a + b", "((!a) + b)")
	test("~a ^ b", "((~a) ^ b)")
	test("-a!", "(-(a!))")
	test("!a!", "(!(a!))")

	// Binary precedence
	test("a = b + c * d ^ e - f / g", "(a = ((b + (c * (d ^ e))) - (f / g)))")

	// Binary associativity
	test("a = b = c", "(a = (b = c))")
	test("a + b - c", "((a + b) - c)")
	test("a * b / c", "((a * b) / c)")
	test("a ^ b ^ c", "(a ^ (b ^ c))")

	// Conditional operator
	test("a ? b : c ? d : e", "(a ? b : (c ? d : e))")
	test("a ? b ? c : d : e", "(a ? (b ? c : d) : e)")
	test("a + b ? c * d : e / f", "((a + b) ? (c * d) : (e / f))")

	// Grouping
	test("a + (b + c) + d", "((a + (b + c)) + d)")
	test("a ^ (b + c)", "(a ^ (b + c))")
	test("(!a)!", "((!a)!)")

	// Show the results
	if failed == 0 {
		fmt.Printf("Passed all %d tests.\n", passed)
	} else {
		fmt.Printf("----\n")
		fmt.Printf("Failed %d out of %d tests.\n", failed, failed+passed)
	}
}

/*
Parses the given chunk of code and verifies that it matches the expected
pretty-printed result.
*/
func test(src string, expected string) {
	l := lexer.New(src)
	p := bantamparser.New(l)

	result, err := p.ParseExpression(0)
	if err != nil {
		failed++
		fmt.Printf("[FAIL] Expected: %s\n", expected)
		fmt.Printf("          Error: %v\n", err)
		return
	}

	var sb strings.Builder
	result.Print(&sb)
	actual := sb.String()

	if expected == actual {
		passed++
	} else {
		failed++
		fmt.Printf("[FAIL] Expected: %s\n", expected)
		fmt.Printf("         Actual: %s\n", actual)
	}
}
