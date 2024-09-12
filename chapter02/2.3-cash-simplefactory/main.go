package main

import (
	cr "cashsimplefactory/cashregister"
	"fmt"
)

func main() {
	price := float64(500)
	num := 50
	discount := 3 // cashType
	// 简单工厂模式根据 discount 的数字选择合适的收费类生成实例
	cash := cr.CreateCashAccept(discount)
	// 通过多态, 可以根据不同收费策略计算得到收费的结果
	totalPrices := cash.AcceptCash(price, num)
	fmt.Printf("totalPrices: %f", totalPrices)
}
