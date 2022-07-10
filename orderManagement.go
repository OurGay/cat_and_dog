
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

func checkSignal(s Signal) error {

	log.WithFields(log.Fields{
		"BuyOpen":   s.BuyOpen,
		"SellOpen":  s.SellOpen,
		"SellClose": s.SellClose,
		"BuyClose":  s.BuyClose,
	}).Info("Signal")

	lbo, lso, lbc, lsc := len(s.BuyOpen), len(s.SellOpen), len(s.BuyClose), len(s.SellClose)

	if lbo > 0 && lbc == 0 {
		return errors.New("Missing BuyClose levels for existing BuyOpen")
	}
