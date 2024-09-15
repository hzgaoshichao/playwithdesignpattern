package pursuitproxy

type Proxy struct {
	pursuit *Pursuit
}

func NewProxy(sg *SchoolGirl) *Proxy {
	return &Proxy{
		pursuit: newPursuit(sg),
	}
}
func (p *Proxy) GiveDolls() {
	p.pursuit.giveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.pursuit.giveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.pursuit.giveChocolate()
}
