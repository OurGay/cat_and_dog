
package mock

import (
	"math/rand"
	"strconv"

	"github.com/apex/log"
	"github.com/santacruz123/trader"
)

type exchng struct {
	pos []trader.Position
	ord []trader.Order
}

//New constructor
func New() trader.Exchanger {
	return &exchng{}
}

func (e *exchng) FetchOrders() (o []trader.Order, err error) {
	log.Info("Fetching orders")
	return
}

func (e *exchng) FetchPositions() (p []trader.Position, err error) {
	log.Info("Fetching positions")
	return
}

func (e *exchng) Orders() (o []trader.Order) {
	return e.ord
}
