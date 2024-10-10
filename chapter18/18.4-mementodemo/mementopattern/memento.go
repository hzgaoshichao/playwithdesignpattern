package mementopattern

type memento struct {
	state string
}

func NewMemento(state string) *memento {
	return &memento{
		state: state,
	}
}

func (m *memento) GetState() string {
	return m.state
}

func (m *memento) SetState(value string) {
	m.state = value
}
