package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("INPUT 1", "telkom")
	fmt.Println("INPUT 2", "tlkom")
	isocc := isOnceChangesChar("telkom", "tlkom")
	fmt.Println("RESULT =>", isocc)
}

func isOnceChangesChar(str1 string, str2 string) bool {
	if str1 == "" || str2 == "" {
		return false
	}

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if str1 == str2 {
		return true
	}

	a1 := []string{}
	b1 := []string{}

	for _, rune1 := range str1 {
		a1 = append(a1, string(rune1))
	}

	for _, rune2 := range str2 {
		b1 = append(b1, string(rune2))
	}

	isSameLen := 0
	if len(a1) < len(b1) {
		a1 = b1
		b1 = a1
	} else if len(a1) == len(b1) {
		a1 = b1
		b1 = a1
		isSameLen++
	}

	lenDiff := len(a1) - len(b1)

	if lenDiff > 1 {
		return false
	}

	n := 0
	j := 0

	cntNotSame := 0
	for n < len(a1) {
		if a1[n] != b1[j] {
			n = n + 1
			cntNotSame++
			continue
		}
		j++
		n++
	}

	if isSameLen > 0 {
		cntNotSame++
	}

	return cntNotSame == 1
}
