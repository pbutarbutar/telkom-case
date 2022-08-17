package utils

import (
	"math/rand"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/google/uuid"
)

type Utility struct{}

func NewUtility() Utility {
	return Utility{}
}

func (u Utility) GenerateRandom(length int) string {
	rand.Seed(time.Now().UnixNano())

	uidGen := uuid.New().String()
	chars := strings.ReplaceAll(uidGen, "-", "")
	var sb strings.Builder
	l := len(chars)

	for i := 0; i < length; i++ {
		c := chars[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return strings.ToUpper(sb.String())
}
func (u Utility) FormatMoneyRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return stringValue
}
