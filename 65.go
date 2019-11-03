package leetcode

// 2019/11/03 21:26 by fzls
// \s* [+-]? (\d+ (\.\d+)? || \.\d+)(e[+-]?\d+)? \s*
// exp: "   -12.345e-678    "
// exp: "   -.345e-678    "
func isNumber(s string) bool {
	return NewNumberValidator(s).IsNumber()
}

type NumberValidator struct {
	s   string
	pos int
}

func NewNumberValidator(s string) *NumberValidator {
	return &NumberValidator{
		s: s,
	}
}

func (v *NumberValidator) IsNumber() bool {
	if v.isEmpty() {
		return false
	}

	v.readOptionalSpace()
	v.readOptionalSign()
	if ok := v.readDecimal(); !ok {
		return false
	}
	if ok := v.readOptionalExponent(); !ok {
		return false
	}
	v.readOptionalSpace()

	return v.isValidFinalState()
}

func (v *NumberValidator) readOptionalSpace() {
	for v.peek() == ' ' {
		v.next()
	}
}

func (v *NumberValidator) readOptionalSign() {
	if v.peek() == '+' || v.peek() == '-' {
		v.next()
	}
}

func (v *NumberValidator) readDecimal() bool {
	// xxx.yyy : if . exists, xxx or yyy must at least exist 1 that is xxx., .yyy, xxx.yyy is ok, but . is not valid
	// .yyy
	if v.peek() == '.' {
		v.next()
		return v.readDigits()
	}

	// xxx.yyy and xxx.
	if ok := v.readDigits(); !ok {
		return false
	}

	if v.peek() == '.' {
		v.next()
		v.readOptionalDigits()
	}

	return true
}

func (v *NumberValidator) readDigits() bool {
	if !v.hasMore() {
		return false
	}

	if !(v.peek() >= '0' && v.peek() <= '9') {
		return false
	}
	v.next()

	v.readOptionalDigits()
	return true
}

func (v *NumberValidator) readOptionalDigits() {
	for v.peek() >= '0' && v.peek() <= '9' {
		v.next()
	}
}

func (v *NumberValidator) readOptionalExponent() bool {
	if !v.hasMore() {
		return true
	}

	if v.peek() != 'e' {
		return true
	}
	// read 'e'
	v.next()

	v.readOptionalSign()
	return v.readDigits()
}

func (v *NumberValidator) isEmpty() bool {
	return len(v.s) == 0
}

func (v *NumberValidator) hasMore() bool {
	return v.pos < len(v.s)
}

func (v *NumberValidator) isValidFinalState() bool {
	return v.pos == len(v.s)
}

func (v *NumberValidator) peek() uint8 {
	if !v.hasMore() {
		return 0
	}

	return v.s[v.pos]
}

func (v *NumberValidator) next() {
	v.pos++
}
