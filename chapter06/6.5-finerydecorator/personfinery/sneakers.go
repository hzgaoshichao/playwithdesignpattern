package personfinery

import "fmt"

type Sneakers struct {
	Finery
}

func (t *Sneakers) Show() {
	fmt.Printf("球鞋 ")
	t.Finery.Show()
}
