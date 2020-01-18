
package trader

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Engine", func() {

	Context("Symbol", func() {
		It("Add", func() {
			ngn := NewEngine()
			quotes := make(chan Quote)
			trades := make(chan Trade)
			ngn.AddSymbol("EUR/USD", quotes, trades)
		})

		It("OHLC", func() {

			symbol := "EUR/USD"

			ngn := NewEngine()
			quotes := make(chan Quote)
			trades := make(chan Trade)
			ngn.AddSymbol(symbol, quotes, trades)
			ngn.Run()
			defer ngn.Stop()