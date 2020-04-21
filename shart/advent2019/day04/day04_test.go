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

func BenchmarkDay4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAllPasswords(day04Input)
	}
}

var _ = Describe("Day 04", func() {
	Describe("GetAllPasswords", func() {
		Context("with known good", func() {
			It("should return a single valid password", func() {
				inputText := "112233-112233"
				actual := GetAllPasswords(inputText)
				Expect(actual).To(Equal([]int{112233}))
				Expect(len(actual)).To(Equal(1))
			})
		})
		Context("with challenge input", func() {
			It("should give me the answer", func() {
				inputText := day04Input
				actual := GetAllPasswords(inputText)
				Expect(len(actual)).To(Equal(1196))
			})
		})
	})
	Describe("ValidatePassword", func() {
		Context("with known good password", func() {
			It("should return true", func() {
				actual := ValidatePassword(112233)
				Expect(actual).To(Equal(true))
				actual = ValidatePassword(111122)
				Expect(actual).To(Equal(true))
				// double at the front
				actual = ValidatePassword(112345)
				Expect(actual).To(Equal(true))
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
		Context("with a password with a decending digit", func() {
			It("should return false", func() {
				actual := ValidatePassword(123450)
				Expect(actual).To(Equal(false))
			})
		})
		Context("with a password where the double is a part of a larger group", func() {
			It("should return false", func() {
				actual := ValidatePassword(123444)
				Expect(actual).To(Equal(false))
			})
		})
	})
})
