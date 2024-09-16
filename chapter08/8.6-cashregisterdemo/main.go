package main

import (
	cr "cashfactorymethod/cashregister"
	"fmt"
)

func main() {
	var discount int = 5
	var price float64 = 500
	var num int = 4

	cashContext := cr.NewCashContext(discount)
	totalPrices := cashContext.GetResult(price, num)
	fmt.Printf("totalPrices: %f", totalPrices)
}
