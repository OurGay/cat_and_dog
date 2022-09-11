
package trader

//Signal - signal of strategy
type Signal struct {
	Strategy                               *Strategy
	BuyOpen, BuyClose, SellOpen, SellClose []Level
}

//Level for price/size
type Level struct {
	Price float64