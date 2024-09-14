package decoratordesignpattern

import "fmt"

type Decorator interface {
	Component
	SetComponent()
}

type DecoratorCommon struct {
	component Component
}

func (c *DecoratorCommon) SetComponent(component Component) {
	c.component = component
}

func (c *DecoratorCommon) Operation() {
	if c.component != nil {
		c.component.Operation()
	} else {
		fmt.Println("component is nil")
	}
}

type ConcreteDecoratorA struct {
	DecoratorCommon
	addedState string
}

func (c *ConcreteDecoratorA) Operation() {
	c.DecoratorCommon.Operation()

	c.addedState = "具体装饰对象 A 的独有操作"
	fmt.Println(c.addedState)
}

type ConcreteDecoratorB struct {
	DecoratorCommon
}

func (c *ConcreteDecoratorB) Operation() {
	c.DecoratorCommon.Operation()
	c.addedBehavior()

}
func (c *ConcreteDecoratorB) addedBehavior() {
	fmt.Println("具体装饰对象 B 的独有操作")
}
