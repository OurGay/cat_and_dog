package trader

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("History", func() {
	Context("CSV", func() {
		symbol := "BTC/USD"

		It("Loading", func() {

			ngn := NewEngine()
			quotes := make(chan Quote)
			trades := make(chan Trade)
			ngn.AddSymbol(sym