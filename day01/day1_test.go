package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSolve(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day 1")
}

var _ = Describe("Day 1", func() {
	It("should validate part 1", func() {
		expenses := []int{1721, 979, 366, 299, 675, 1456}
		value, err := Part1(expenses)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(514579))
	})

	It("should validate part 2", func() {
		expenses := []int{1721, 979, 366, 299, 675, 1456}
		value, err := Part2(expenses)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(241861950))
	})
})
