package personfinery

import "fmt"

type LeatherShoes struct {
	Finery
}

func (t *LeatherShoes) Show() {
	fmt.Printf("皮鞋 ")
	t.Finery.Show()
}
