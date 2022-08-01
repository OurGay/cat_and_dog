
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
		}

		sort.Sort(byPriceDesc(signal.BuyOpen))
		sort.Sort(byPrice(signal.BuyClose))
		sort.Sort(byPrice(signal.SellOpen))
		sort.Sort(byPriceDesc(signal.SellClose))

		if err := checkSignal(signal); err != nil {
			err := errors.Wrap(err, "Signal error")
			log.Error(err.Error())
			continue
		}

		o.lastSignal = signal

		o.openOrders(true, signal.BuyOpen)
		o.openOrders(false, signal.SellOpen)
		o.closeOrders(true, signal.BuyClose)
		o.closeOrders(false, signal.SellClose)

		o.diffOrders()

		o.sendOrders()
	}
}

func (o *orderManagement) openOrders(buySell bool, levels []Level) {
	var sign int

	if buySell {
		sign = 1
	} else {
		sign = -1
	}

	position := o.position(o.strategy.Symbol).Amount

	for _, level := range levels {
		amount := sign * level.Size * o.strategy.Size

		if position != 0 && ((buySell && position-amount > 0) || (!buySell && position-amount < 0)) {
			position, amount = position-amount, 0
		} else {
			position, amount = 0, -1*(position-amount)
		}

		if amount != 0 {
			o.newOrders = append(o.newOrders,
				Order{
					Symbol: o.strategy.Symbol,
					Price:  level.Price,
					Amount: amount,
					Type:   Limit,
				},
			)
		}
	}
}

func (o *orderManagement) closeOrders(buySell bool, levels []Level) {
	position := o.position(o.strategy.Symbol).Amount
	if position == 0 || (buySell && position < 0) || (!buySell && position > 0) {
		return
	}

	var sign int

	if buySell {
		sign = -1
	} else {
		sign = 1
	}

	for _, level := range levels {
		if position == 0 {
			return
		}

		amount := sign * level.Size * o.strategy.Size

		if (buySell && position+amount > 0) || (!buySell && position+amount < 0) {
			position += amount
		} else {
			position, amount = 0, -position
		}