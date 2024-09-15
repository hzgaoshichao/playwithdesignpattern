package main

import px "proxydesignpattern/pursuitproxy"

func main() {
	sg := px.SchoolGirl{}
	sg.SetName("小美")

	p := px.NewProxy(&sg)
	p.GiveDolls()
	p.GiveFlowers()
	p.GiveChocolate()

}
