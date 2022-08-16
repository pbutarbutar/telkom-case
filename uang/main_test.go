package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOfMoney(t *testing.T) {
	expected1 := []string{"'Rp. 100000': 1", "'Rp. 20000': 2", "'Rp. 5000': 1"}
	result := countOfMoney(145000)
	k := 0
	for _, cntMoney := range result {
		assert.Equal(t, cntMoney, expected1[k])
		k++
	}

	expected2 := []string{"'Rp. 2000': 1", "'Rp. 100': 1"}
	result2 := countOfMoney(2050)
	k = 0
	for _, cntMoney := range result2 {
		assert.Equal(t, cntMoney, expected2[k])
		k++
	}

	var expected3 []string
	result3 := countOfMoney(0)
	assert.Equal(t, result3, expected3)
}
