package decoratordesignpattern

import "fmt"

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() {
	fmt.Println("具体对象的实际操作")
}
