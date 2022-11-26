package internal

import (
	"strings"
	"unicode"
)

var (
	MinLenKW  = 3
	MaxLenKW  = 15
	DigitKW   = false
	UpperKW   = false
	LowerKW   = false
	SpecialKW = false
	KW        = ""
	KWForCompare = ""
)

func CheckForDigits(keyWord string) bool {
	for _, c := range keyWord {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func CheckForUppercase(keyWord string) bool {
	for _, c := range keyWord {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func CheckForLowercase(keyWord string) bool {
	for _, c := range keyWord {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false
}

func CheckForSpecialSymbols(keyWord string) bool {
	s := "! ? # $ & % @ ^ * whitespace [ ] ( ) { } < > = ~ | _ ' \" \\ / : ; + - , . `"
	for _, c := range keyWord {
		if strings.Contains(s, string(c)) {
			return true
		}
	}
	return false
}

func CheckKeyWord(keyWord string) bool {
	if DigitKW {
		if !CheckForDigits(keyWord) {
			return false
		}
	}
	if UpperKW {
		if !CheckForUppercase(keyWord) {
			return false
		}
	}
	if LowerKW {
		if !CheckForLowercase(keyWord) {
			return false
		}
	}
	if SpecialKW {
		if !CheckForSpecialSymbols(keyWord) {
			return false
		}
	}

	return len([]rune(keyWord)) >= MinLenKW && len([]rune(keyWord)) <= MaxLenKW
}
