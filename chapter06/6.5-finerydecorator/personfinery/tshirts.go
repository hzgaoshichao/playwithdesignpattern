package personfinery

import "fmt"

type Tshirts struct {
	Finery
}

func (t *Tshirts) Show() {
	fmt.Printf("大T恤 ")
	t.Finery.Show()
}
