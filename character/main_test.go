package main

import (
	"testing"
)

func TestIsOnceChangesChar(t *testing.T) {

	isTrue1 := isOnceChangesChar("telkom", "tlkom")
	if isTrue1 != true {
		t.Errorf("Incorrect! should %t", true)
	}

	isTrue2 := isOnceChangesChar("telkom", "telkom")
	if isTrue2 != true {
		t.Errorf("Incorrect! should %t", true)
	}

	isFalse1 := isOnceChangesChar("telkom", "telecom")
	if isFalse1 == true {
		t.Errorf("Incorrect! should %t", false)
	}

	isFalse2 := isOnceChangesChar("tllkom", "tlkmome")
	if isFalse2 == true {
		t.Errorf("Incorrect! should %t", false)
	}

	isFalse3 := isOnceChangesChar("", "")
	if isFalse3 == true {
		t.Errorf("Incorrect! should %t", false)
	}

	isFalse4 := isOnceChangesChar("telkom", "")
	if isFalse4 == true {
		t.Errorf("Incorrect! should %t", false)
	}
}
