package personfinery

import "fmt"

type Suit struct {
	Finery
}

func (t *Suit) Show() {
	fmt.Printf("西装 ")
	t.Finery.Show()
}
