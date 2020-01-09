
package trader

import (
	"time"

	"github.com/apex/log"
)

//Engine that runs strategies
type Engine struct {
	strategy *Strategy
	om       *orderManagement

	signals chan Signal

	tradeCh  map[string]chan Trade
	quoteCh  map[string]chan Quote
	changeCh map[string]chan struct{}

	ohlc   map[string]*History
	quotes map[string]*Quote

	quit chan struct{}
}

//Trade type
type Trade struct {
	Price, Amount float64
	Time          time.Time
}

//Quote type
type Quote struct {
	//TODO refactor into int
	Bid, Ask float64
}

//NewEngine constructor
func NewEngine() *Engine {
	return &Engine{
		quoteCh:  make(map[string]chan Quote),
		tradeCh:  make(map[string]chan Trade),
		changeCh: make(map[string]chan struct{}),
		ohlc:     make(map[string]*History),
		quit:     make(chan struct{}),
		signals:  make(chan Signal),
	}
}

//AddSymbol to engine - quotes and trades for this symbol
func (e *Engine) AddSymbol(symbol string, quotes chan Quote, trades chan Trade) {
	e.quoteCh[symbol] = quotes
	e.tradeCh[symbol] = trades
	e.changeCh[symbol] = make(chan struct{})

	initTimeSeries(e, symbol)

	e.quotes = make(map[string]*Quote)
	e.quotes[symbol] = &Quote{}
}

//AddStrategy to engine
func (e *Engine) AddStrategy(strategy *Strategy) {