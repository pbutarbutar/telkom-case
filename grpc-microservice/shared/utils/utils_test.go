package utils

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utility Test", func() {
	BeforeEach(func() {

	})

	It("GenerateRandom Test ", func() {
		uT := Utility{}
		newU := NewUtility()
		Expect(reflect.TypeOf(newU)).To(Equal(reflect.TypeOf(uT)))

		gen := uT.GenerateRandom(8)
		Expect(gen).NotTo(Equal(""))
		Expect(len(gen)).To(Equal(8))

		genErr := uT.GenerateRandom(0)
		Expect(genErr).To(Equal(""))
	})

})
