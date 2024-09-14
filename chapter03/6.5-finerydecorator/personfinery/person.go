package personfinery

import "fmt"

type Person struct {
	name string
}

func NewPerson(name string) Person {
	return Person{name: name}
}

func (p *Person) Show() {
	fmt.Printf("装扮的 %s \n", p.name)
}
