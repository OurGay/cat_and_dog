
package trader

//Order abstration
type Order struct {
	ID     string
	Symbol string
	Price  float64
	Amount int
	Type   OrderType
}

//Position abstraction