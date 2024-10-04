package companyobserver

import "fmt"

type NBAObserver struct {
	Observer
}

func (n NBAObserver) update() {
	fmt.Printf("%v: %v %v 请关闭股票行情,赶紧离开 \n", n.sub.GetSubject().name, n.sub.GetSubject().GetAction(), n.Observer.name)
}

func NewNBAObserver(name string, sub SubjectInterface) NBAObserver {
	return NBAObserver{
		Observer{name: name, sub: sub},
	}
}
