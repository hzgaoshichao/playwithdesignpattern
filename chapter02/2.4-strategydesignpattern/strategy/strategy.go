package strategy

import "fmt"

type Strategy interface {
	algorithmInterface()
}

type ConcreteStrategyA struct {
}

func (cs *ConcreteStrategyA) algorithmInterface() {
	fmt.Println("算法 A 实现")
}

type ConcreteStrategyB struct {
}

func (cs *ConcreteStrategyB) algorithmInterface() {
	fmt.Println("算法 B 实现")
}

type ConcreteStrategyC struct {
}

func (cs *ConcreteStrategyC) algorithmInterface() {
	fmt.Println("算法 C 实现")
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
	}

}

func (c *Context) ContextInterface() {
	c.strategy.algorithmInterface()
}
