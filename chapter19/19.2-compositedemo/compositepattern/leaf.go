package compositepattern

import "fmt"

type leaf struct {
	component
}

func NewLeaf(name string) *leaf {
	return &leaf{
		component: component{name: name},
	}
}

func (l *leaf) Add(component componentInterface) {
	fmt.Printf("Can not add to a leaf \n")
}

func (l *leaf) Remove(component componentInterface) {
	fmt.Printf("Can not remove to a leaf \n")
}

func (l *leaf) Display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}

	fmt.Printf("%v \n", l.component.name)

}
