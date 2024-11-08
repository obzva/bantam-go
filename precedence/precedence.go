package precedence

type Precedence int

const (
	ASSIGNMENT Precedence = iota + 1
	CONDITIONAL
	SUM
	PRODUCT
	EXPONENT
	PREFIX
	POSTFIX
	CALL
)
