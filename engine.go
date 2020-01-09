
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