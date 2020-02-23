package bitfinex

import (
	"math/rand"
	"strconv"

	"github.com/apex/log"
	"github.com/santacruz123/trader"
)

type exchng struct {
	pos []trader.Position
	ord []trader.Or