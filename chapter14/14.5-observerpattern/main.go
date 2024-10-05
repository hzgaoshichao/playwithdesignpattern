package main

import odp "observerpattern/observerdp"

func main() {

	subject := odp.NewConcreteSubject()

	nameX := odp.NewConcreteObserver("NameX", subject)
	nameY := odp.NewConcreteObserver("NameY", subject)
	nameZ := odp.NewConcreteObserver("NameZ", subject)
	subject.Attach(nameX)
	subject.Attach(nameY)
	subject.Attach(nameZ)

	subject.SetSubjectState("ABC")
	subject.NotifyObserver()

}
