
package trader

import (
	"reflect"
	"sort"

	"github.com/apex/log"
	"github.com/pkg/errors"
)

type orderManagement struct {
	strategy *Strategy

	lastSignal Signal
	signalCh   chan Signal

	sendOrdersOn            bool
	newOrders, cancelOrders []Order

	exchange Exchanger
}

func newOrderManagement(s *Strategy, e Exchanger) *orderManagement {
	return &orderManagement{
		sendOrdersOn: true,
		strategy:     s,
		exchange:     e,
		signalCh:     make(chan Signal),
	}
}
