package mementopattern

import "fmt"

type Origator struct {
	state string
}

func (o *Origator) SetState(value string) {
	o.state = value
}

func (o *Origator) Show() {
	fmt.Printf("State %v \n", o.state)
}

func (o *Origator) CreateMemento() *memento {
	return &memento{state: o.state}
}

func (o *Origator) RecoverMemento(m *memento) {
	o.state = m.GetState()
}
