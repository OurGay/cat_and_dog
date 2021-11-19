
package trader

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

func initTimeSeries(e *Engine, s string) {
	h := &History{
		TimeSeries: make(map[Timeframe]map[time.Time]*OHLC),
	}

	e.ohlc[s] = h

	h.TimeSeries = map[Timeframe]map[time.Time]*OHLC{