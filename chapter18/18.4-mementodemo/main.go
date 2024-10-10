package main

import mp "mementodemo/mementopattern"

func main() {
	o := mp.Origator{}
	o.SetState("On")
	o.Show()

	c := mp.Caretaker{}
	c.SetMemento(o.CreateMemento())

	o.SetState("OFF")
	o.Show()

	o.RecoverMemento(c.GetMemento())
	o.Show()
}
