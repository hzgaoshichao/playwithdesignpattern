package mementopattern

type Caretaker struct {
	memento *memento
}

func (c *Caretaker) GetMemento() *memento {
	return c.memento
}

func (c *Caretaker) SetMemento(value *memento) {
	c.memento = value
}
