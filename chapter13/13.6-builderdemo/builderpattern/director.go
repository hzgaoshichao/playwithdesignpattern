package builderpattern

// 指挥者
type Director struct {
}

func (d *Director) Construct(builder Builder) {
	builder.buildPartA()
	builder.buildPartB()
}
