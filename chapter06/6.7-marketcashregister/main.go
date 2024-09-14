package main

import (
	cr "cashsimplefactorydecorator/cashregister"
	"fmt"
)

func main() {
	var discount int = 6    //折扣模式
	var price float64 = 500 // 商品单价
	var num int = 4         //商品数量

	cashContext := cr.NewCashContext(discount)
	totalPrices := cashContext.GetResult(price, num)
	fmt.Printf("totalPrices: %f \n", totalPrices)
}
