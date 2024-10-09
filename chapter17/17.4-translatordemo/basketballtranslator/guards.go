package basketballtranslator

import "fmt"

type guards struct {
	player
}

func NewGuards(name string) *guards {
	return &guards{
		player{
			name: name,
		},
	}
}

func (g *guards) Attack() {
	fmt.Printf("guards attack!!!! \n")
}

func (g *guards) Defense() {
	fmt.Printf("guards defense!!!! \n")
}
