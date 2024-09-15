package pursuitproxy

type Proxy struct {
	// 这里使用抽象接口来代替原来具体的结构体，这样可以提升可扩展性
	// pursuit  *Pursuit
	giveGift GiveGift
}

func NewProxy(sg *SchoolGirl) *Proxy {
	return &Proxy{
		// pursuit: newPursuit(sg),
		giveGift: newPursuit(sg),
	}
}
func (p *Proxy) GiveDolls() {
	// p.pursuit.GiveDolls()
	p.giveGift.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	// p.pursuit.GiveFlowers()
	p.giveGift.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	// p.pursuit.GiveChocolate()
	p.giveGift.GiveChocolate()
}
