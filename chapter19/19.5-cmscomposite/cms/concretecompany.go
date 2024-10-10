package cms

import (
	"fmt"
	"reflect"
)

type concreteCompany struct {
	name     string
	children []companyInterface
}

func NewConcreteCompany(name string) *concreteCompany {
	return &concreteCompany{
		name:     name,
		children: make([]companyInterface, 0),
	}
}

func (c *concreteCompany) Add(com companyInterface) {
	c.children = append(c.children, com)
}

func (c *concreteCompany) Remove(com companyInterface) {
	for i, item := range c.children {
		if reflect.DeepEqual(item, com) {
			c.children = append(c.children[:i], c.children[i+1:]...)
		}
	}
}

func (c *concreteCompany) Display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%v \n", c.name)

	for _, item := range c.children {
		item.Display(depth + 2)
	}
}

func (c *concreteCompany) LineOfDuty() {

	for _, item := range c.children {
		item.LineOfDuty()
	}
}
