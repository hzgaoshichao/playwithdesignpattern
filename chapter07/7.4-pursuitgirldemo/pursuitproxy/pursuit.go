package pursuitproxy

import "fmt"

type Pursuit struct {
	schoolgirl *SchoolGirl
}

func newPursuit(sg *SchoolGirl) *Pursuit {
	return &Pursuit{
		schoolgirl: sg,
	}
}

func (p *Pursuit) GiveDolls() {
	fmt.Printf("%v, 你好, 送你洋娃娃 \n", p.schoolgirl.name)
}

func (p *Pursuit) GiveFlowers() {
	fmt.Printf("%v, 你好, 送你鲜花 \n", p.schoolgirl.name)
}

func (p *Pursuit) GiveChocolate() {
	fmt.Printf("%v, 你好, 送你巧克力 \n", p.schoolgirl.name)
}
