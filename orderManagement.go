
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

	if lso > 0 && lsc == 0 {
		return errors.New("Missing SellClose levels for existing SellOpen")
	}

	if lbc > 0 && lso > 0 && s.BuyClose[lbc-1].Price > s.SellOpen[0].Price {
		return errors.New("BuyClose overlaps SellOpen")
	}

	if lsc > 0 && lbo > 0 && s.SellClose[lsc-1].Price > s.BuyOpen[0].Price {
		return errors.New("SellClose overlaps BuyOpen")
	}

	if lsc > 0 && lso > 0 && s.SellOpen[0].Price > s.SellClose[lso-1].Price {
		return errors.New("SellClose overlaps SellOpen")
	}

	if lbc > 0 && lbo > 0 && s.BuyOpen[0].Price > s.BuyClose[0].Price {
		return errors.New("BuyClose overlaps BuyOpen")
	}

	return nil
}

func (o *orderManagement) signalLoop() {
	for {
		signal := <-o.signalCh

		if reflect.DeepEqual(o.lastSignal, signal) {
			continue