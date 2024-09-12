package main

import (
	"fmt"
	cr "strategysimplefactory/cashregister"
)

func main() {
	var price float64 = 500
	var num int = 50
	var discount int = 3
	// 根据用户输入,将对于的策略对象作为参数传入 CashContext 对象中
	cashContext := cr.NewCashContext(discount)
	totalPrices := cashContext.GetResult(price, num)
	fmt.Printf("totalPrices: %f", totalPrices)
}
