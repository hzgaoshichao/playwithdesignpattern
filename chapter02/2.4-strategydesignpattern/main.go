package main

import s "strategydesignpattern/strategy"

func main() {
	var context *s.Context

	context = s.NewContext(&s.ConcreteStrategyA{})
	context.ContextInterface()

	context = s.NewContext(&s.ConcreteStrategyB{})
	context.ContextInterface()

	context = s.NewContext(&s.ConcreteStrategyC{})
	context.ContextInterface()
}
