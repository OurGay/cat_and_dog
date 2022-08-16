
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