package bitfinex

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

//NewExchanger constructor
func New() trader.Exchanger {
	e := &exchng{}

	go func() {}()

	return e
}

func (e *exchng) FetchOrders