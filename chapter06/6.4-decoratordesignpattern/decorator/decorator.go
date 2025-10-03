package decoratordesignpattern

import "fmt"

// 是否可以去掉 Decorator 接口，取决于是否需要“统一约束装饰器类型”这个抽象。
// 情况说明：
// - 当前 Decorator 接口本身没有被使用（若项目中没有引用它），属冗余，可删除。
// - 你的具体装饰器（ConcreteDecoratorA / B）已经通过嵌入 DecoratorCommon 实现 Component 行为。
// - 只有当你需要：
//  1. 统一接收“任意装饰器”参数（如 func Wrap(d Decorator)）
//  2. 做运行时类型断言
//  3. 对外屏蔽内部实现（只暴露装饰器能力）
//     时，才保留 Decorator 接口有价值。
type Decorator interface {
	Component
	SetComponent(c Component)
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
