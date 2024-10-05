package observerdp

import "fmt"

type Observer interface {
	update()
}

type ConcreteObserver struct {
	name string
	sub  Subject
}

func (co *ConcreteObserver) update() {
	fmt.Printf("观察者%v的新状态是%v \n", co.name, co.sub.GetSubjectState())
}

func NewConcreteObserver(name string, sub Subject) *ConcreteObserver {
	return &ConcreteObserver{
		name: name,
		sub:  sub,
	}
}
