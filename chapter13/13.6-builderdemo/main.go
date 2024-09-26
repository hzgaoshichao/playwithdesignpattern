package main

import bp "builderdemo/builderpattern"

func main() {
	director := bp.Director{}
	b1 := bp.NewConcreteBuilder1()

	// 指挥者用ConcreteBuilder1的方法来建造产品
	director.Construct(b1) // 创建的是产品 由部件 A 和部件 B 组成
	p1 := b1.GetResult()
	p1.Show()

	// 指挥者用ConcreteBuilder2的方法来建造产品
	b2 := bp.NewConcreteBuilder2() // 创建的是产品 由部件 X 和部件 Y 组成
	director.Construct(b2)
	p2 := b2.GetResult()
	p2.Show()
}

// 这个例子更加的容易理解, 更加的优雅： https://github.com/ruanrunxue/Practice-Design-Pattern--Go-Implementation/blob/main/docs/go_practice_design_pattern__builder.md
