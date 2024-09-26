package builderpattern

// 具体建造者1
type ConcreteBuilder1 struct {
	product *product
}

func (c *ConcreteBuilder1) buildPartA() {
	c.product.add("部件 A")
}

func (c *ConcreteBuilder1) buildPartB() {
	c.product.add("部件 B")
}

func (c *ConcreteBuilder1) GetResult() *product {
	return c.product
}

func NewConcreteBuilder1() *ConcreteBuilder1 {
	return &ConcreteBuilder1{
		product: &product{parts: make([]string, 0)},
	}
}
