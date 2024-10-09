package main

import sp "statepatterndemo/statepattern"

func main() {
	c := sp.NewContext(&sp.ConcreteStateA{})
	c.Request()
	c.Request()
	c.Request()
	c.Request()

}
