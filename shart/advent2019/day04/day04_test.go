package day04_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sean-hart/advent2019/shart/advent2019/day04"
)

var day04Input string = "152085-670283"

func TestDay04(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day 04")
}

var _ = Describe("Day 04", func() {
	Describe("GetAllPasswords", func() {
		Context("with known good", func() {
			It("should return a single valid password", func() {
				inputText := "111111-111111"
				actual := GetAllPasswords(inputText)
				Expect(actual).To(Equal([]int{111111}))
				Expect(len(actual)).To(Equal(1))
			})
		})
		Context("with challenge input", func() {
			It("should give me the answer", func() {
				inputText := day04Input
				actual := GetAllPasswords(inputText)
				Expect(len(actual)).To(Equal(1764))
			})
		})
	})
	Describe("ValidatePassword", func() {
		Context("with known good password", func() {
			It("should return true", func() {
				By("Being All ones", func() {
					actual := ValidatePassword(111111)
					Expect(actual).To(Equal(true))
				})
				By("Being a two in the final spot", func() {
					actual := ValidatePassword(111112)
					Expect(actual).To(Equal(true))
				})
			})
		})
		Context("with a password with no doubles", func() {
			It("should return false", func() {
				actual := ValidatePassword(123456)
				Expect(actual).To(Equal(false))
			})
		})
		Context("with a password that is not 6 characters", func() {
			It("should return false", func() {
				actual := ValidatePassword(11111)
				Expect(actual).To(Equal(false))
				actual = ValidatePassword(1111111)
				Expect(actual).To(Equal(false))
			})
		})
		Context("with a password that goes down after first digit", func() {
			It("should return false", func() {
				actual := ValidatePassword(123450)
				Expect(actual).To(Equal(false))
			})
		})
	})
})
