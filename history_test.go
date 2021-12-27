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
			ngn.AddSymbol(symbol, quotes, trades)

			f, err := os.Open("fixtures/m30.csv")
			Expect(err).To(Succeed())

			err = ngn.LoadHistory(symbol, M30, f)
			Expect(err).To(Succeed())

			t, err := time.Parse("2006-01-02 15:04:05", "2017-