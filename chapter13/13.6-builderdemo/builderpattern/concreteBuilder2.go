package builderpattern

// 具体建造者2
type ConcreteBuilder2 struct {
	product *product
}

func (c *ConcreteBuilder2) buildPartA() {
	c.product.add("部件 X")
}

func (c *ConcreteBuilder2) buildPartB() {
	c.product.add("部件 Y")
}

func (c *ConcreteBuilder2) GetResult() *product {
	return c.product
}

func NewConcreteBuilder2() *ConcreteBuilder2 {
	return &ConcreteBuilder2{
		product: &product{parts: make([]string, 0)},
	}
}
