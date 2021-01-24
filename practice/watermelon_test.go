package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "practice"
)

var _ = Describe("Basic test", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(Divide(4)).To(Equal(true))
		Expect(Divide(2)).To(Equal(false))
		Expect(Divide(5)).To(Equal(false))
		Expect(Divide(88)).To(Equal(true))
		Expect(Divide(100)).To(Equal(true))
		Expect(Divide(67)).To(Equal(false))
		Expect(Divide(90)).To(Equal(true))
		Expect(Divide(10)).To(Equal(true))
		Expect(Divide(99)).To(Equal(false))
		Expect(Divide(32)).To(Equal(true))
	})
})
