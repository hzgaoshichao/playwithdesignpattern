package personfinery

import "fmt"

type BigTrouser struct {
	Finery
}

func (b *BigTrouser) Show() {
	fmt.Printf("垮裤 ")
	b.Finery.Show()
}
