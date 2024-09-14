package decoratordesignpattern

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() {
	fmt.Println("具体对象的实际操作")
}
