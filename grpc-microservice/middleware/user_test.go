package middleware

import (
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo/v2"
)

var user, humanid, email string

var _ = Describe("Middleware User", func() {

	BeforeEach(func() {

	})

	Describe("Test User ", func() {
		It("Clone", func() {
			user = "UserID"
			humanid = "HumanID"
			email = "Email@email.com"
			nUser, err := NewUser(user, humanid, email)
			Expect(err).To(BeNil())

			cln := nUser.Clone()
			Expect(cln.UserID).To(Equal(user))
			Expect(cln.HumanID).To(Equal(humanid))
			Expect(cln.Email).To(Equal(email))

		})
	})

})
