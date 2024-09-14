package personfinery

import "fmt"

type Strawhat struct {
	Finery
}

func (t *Strawhat) Show() {
	fmt.Printf("草帽 ")
	t.Finery.Show()
}
