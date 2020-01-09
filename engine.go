
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