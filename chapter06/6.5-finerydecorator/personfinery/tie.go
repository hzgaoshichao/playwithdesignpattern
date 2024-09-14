package personfinery

import "fmt"

type Tie struct {
	Finery
}

func (t *Tie) Show() {
	fmt.Printf("领带 ")
	t.Finery.Show()
}
