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

func isOnceChangesChar(char1 string, char2 string) bool {

	if char1 == "" || char2 == "" {
		return false
	}

	char1 = strings.ToLower(char1)
	char2 = strings.ToLower(char2)

	if char1 == char2 {
		return true
	}

	lenDiff := len(char1) - len(char2)

	if lenDiff < 0 {
		lenDiff = lenDiff * -1
	}
	if lenDiff > 1 {
		return false
	}

	a := []string{}
	b := []string{}

	for _, rune1 := range char1 {
		a = append(a, string(rune1))
	}

	for _, rune2 := range char2 {
		b = append(b, string(rune2))
	}

	a1 := []string{}
	b1 := []string{}

	isSameLen := 0
	if len(a) > len(b) {
		a1 = a
		b1 = b
	} else if len(a) == len(b) {
		a1 = a
		b1 = b
		isSameLen++
	} else {
		a1 = b
		b1 = a
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
