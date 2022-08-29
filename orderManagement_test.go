
package trader

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type exchng struct {
	pos []Position
	ord []Order
}

func (e *exchng) FetchOrders() (o []Order, err error) {
	return
}

func (e *exchng) FetchPositions() (p []Position, err error) {
	return
}

func (e *exchng) Orders() (o []Order) {
	return e.ord
}

func (e *exchng) Positions() (p []Position) {
	return e.pos
}

func (e *exchng) NewOrder(o Order) (ID string, err error) {
	return
}

func (e *exchng) CancelOrder(ID string) (ok bool, err error) {
	return
}

var _ = Describe("OrderManagement", func() {

	symbol := "BTC/USD"

	strategy := &Strategy{
		Title:  "Some",
		Code:   "qqq",
		Symbol: symbol,
		Size:   100,
		Parts:  5,
	}

	Context("Signal validity", func() {
		It("SellOpen overlaps with BuyClose", func() {
			signal := Signal{
				BuyClose: []Level{
					Level{101, 1},
					Level{102, 1},
					Level{106, 2},
					Level{104, 1},
				},
				SellOpen: []Level{
					Level{103, 1},
				},
			}
