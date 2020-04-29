package libs_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/sean-hart/advent2019/shart/libs"
)

func TestLibs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Libs")
}

var _ = Describe("getDigits", func() {
	DescribeTable("Good outputs",
		func(inputInt int, expected []int) {
			digits := libs.GetDigits(inputInt)
			Expect(digits).To(Equal(expected))
		},
		Entry("Under 10 - 5 should 0 fill", 5, []int{5, 0, 0, 0, 0}),
		Entry("under 100 - 15 should 0 fill", 15, []int{5, 1, 0, 0, 0}),
		Entry("under 1000 - 125 should 0 fill", 125, []int{5, 2, 1, 0, 0}),
		Entry("Over 10000 - 123456789 should not 0 fill", 123456789, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}),
	)
})

func BenchmarkDay5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 2
		max := int(math.Pow(10, 9))
		randomNumber := rand.Intn(max-min) + min

		libs.GetDigits(randomNumber)
	}
}
