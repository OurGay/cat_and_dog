package trader

import (
	"time"
)

//Timeframe type
type Timeframe uint8

//Timeframes
const (
	M1 Timeframe = iota
	M5
	M15
	