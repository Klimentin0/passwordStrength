package main

import "unicode"

func isValidPassword(password string) bool {
	str := password
	var sym []string
	for _, c := range str {
		sym = append(sym, string(c))
	}
	hasUpper := false
	hasNumber := false
	for _, letter := range sym {
		if unicode.IsUpper([]rune(letter)[0]) {
			hasUpper = true
		}
		if unicode.IsDigit([]rune(letter)[0]) {
			hasNumber = true
		}

		if hasNumber && hasUpper && 5 <= len(sym) && len(sym) <= 12 {
			return true
		}
	}
	return false
}
