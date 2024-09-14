package main

import d "decoratordesignpattern/decorator"

func main() {
	c := d.ConcreteComponent{}
	d1 := d.ConcreteDecoratorA{}
	d2 := d.ConcreteDecoratorB{}

	d1.SetComponent(&c)
	d2.SetComponent(&d1)
	d2.Operation()

}
