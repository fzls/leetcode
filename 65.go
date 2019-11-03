package leetcode

// 2019/11/03 21:26 by fzls
// \s* [+-]? (\d+ (\.\d+)? || \.\d+)(e[+-]?\d+)? \s*
// exp: "   -12.345e-678    "
// exp: "   -.345e-678    "
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	pos := 0
	ok := true

	readOptionalSpace(s, &pos)
	readOptionalSign(s, &pos)
	if ok = readDecimal(s, &pos); !ok {
		return false
	}
	if ok = readOptionalExponent(s, &pos); !ok {
		return false
	}
	readOptionalSpace(s, &pos)

	return pos == len(s)
}

func readOptionalSpace(s string, pos *int) {
	for *pos < len(s) && s[*pos] == ' ' {
		*pos++
	}
}

func readOptionalSign(s string, pos *int) {
	if *pos < len(s) && (s[*pos] == '+' || s[*pos] == '-') {
		*pos++
	}
}

func readDecimal(s string, pos *int) bool {
	// xxx.yyy : if . exists, xxx or yyy must at least exist 1 that is xxx., .yyy, xxx.yyy is ok, but . is not valid
	// .yyy
	if *pos < len(s) && s[*pos] == '.' {
		*pos++
		return readDigits(s, pos, true)
	}

	// xxx.yyy and xxx.
	if ok := readDigits(s, pos, false); !ok {
		return false
	}

	if *pos < len(s) && s[*pos] == '.' {
		*pos++
		readOptionalDigits(s, pos)
	}

	return true
}

func readDigits(s string, pos *int, isDecimalPart bool) bool {
	if *pos >= len(s) {
		return false
	}

	if !(s[*pos] >= '0' && s[*pos] <= '9') {
		return false
	}
	*pos++

	for *pos < len(s) && s[*pos] >= '0' && s[*pos] <= '9' {
		*pos++
	}
	return true
}

func readOptionalDigits(s string, pos *int) {
	for *pos < len(s) && s[*pos] >= '0' && s[*pos] <= '9' {
		*pos++
	}
}

func readOptionalExponent(s string, pos *int) bool {
	if *pos >= len(s) {
		return true
	}

	if s[*pos] != 'e' {
		return true
	}
	// read 'e'
	*pos++

	readOptionalSign(s, pos)
	return readDigits(s, pos, false)
}
