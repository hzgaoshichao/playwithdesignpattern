package main

import (
	"fmt"
	i "iteratordemo/iteratorpattern"
)

func main() {
	bus := i.NewConcreteAggregate()
	bus.Add("大鸟")
	bus.Add("小菜")
	bus.Add("行李")
	bus.Add("老外")
	bus.Add("公交内部员工")
	bus.Add("小偷")

	// 正序迭代器
	// conductor := NewConcreteIterator(bus)

	// 或者使用方法来创建正序迭代器
	// conductor := bus.createIterator()

	// 倒序迭代器
	conductor := i.NewConcreteIteratorDesc(bus)

	conductor.First()
	for {
		if conductor.IsDone() {
			break
		}
		fmt.Printf("%v ,请买票上车 \n", conductor.CurrentItem())
		conductor.Next()
	}
}
