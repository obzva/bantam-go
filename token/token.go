package token

import "bantam/tokentype"

/*
I follow the advice from the book 'Jon Bodner - Learning Go_ An Idiomatic Approach to Real-world Go Programming, 2nd Edition-O'Reilly Media (2024)'
So I didn't write getter and setter method but rather access fields directly.
*/
type Token struct {
	Type tokentype.TokenType
	Text string
}
