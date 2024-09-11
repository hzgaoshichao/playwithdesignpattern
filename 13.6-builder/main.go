package main

func main() {
	director := director{}
	b1 := concreteBuilder1{product: &product{parts: make([]string, 0)}}
	director.construct(&b1)
	p1 := b1.getResult()
	p1.show()

	b2 := concreteBuilder2{product: &product{parts: make([]string, 0)}}
	director.construct(&b2)
	p2 := b2.getResult()
	p2.show()
}

// 这个列子更加的容易理解： https://learnku.com/docs/go-patterns/1.0.0/jian-zao-zhe-mo-shi-builder-pattern/14747
// 这个列子更加的容易理解, 更加的优雅： https://github.com/ruanrunxue/Practice-Design-Pattern--Go-Implementation/blob/main/docs/go_practice_design_pattern__builder.md
