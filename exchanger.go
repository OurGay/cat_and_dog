
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
type Position struct {
	Symbol string
	Price  float64
	Amount int
}

//OrderType for orders
type OrderType uint

//Limit,Stop,Market types
const (