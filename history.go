
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
		M1:  make(map[time.Time]*OHLC),
		M5:  make(map[time.Time]*OHLC),
		M15: make(map[time.Time]*OHLC),
		M30: make(map[time.Time]*OHLC),
		H1:  make(map[time.Time]*OHLC),
		H4:  make(map[time.Time]*OHLC),
		D1:  make(map[time.Time]*OHLC),
	}
}

//LoadHistory of instrument
func (e *Engine) LoadHistory(symbol string, tf Timeframe, reader io.Reader) error {

	const format = "2006-01-02 15:04:05"

	r := csv.NewReader(reader)

	initTimeSeries(e, symbol)

	for {
		rec, err := r.Read()

		if err == io.EOF {
			break