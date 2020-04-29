package libs_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/sean-hart/advent2019/shart/libs"
	"math/rand"
	"math"
	"time"

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
		Entry("5", 5, []int{5}),
		Entry("15", 15, []int{5,1}),
		Entry("125", 125, []int{5,2,1}),
	)
})

func BenchmarkDay5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 2
		max := int(math.Pow(10, 9))
		randomNumber := rand.Intn(max - min) + min
		
		libs.GetDigits(randomNumber)
	}
}
