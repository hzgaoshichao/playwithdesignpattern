package personfinery

type Finery struct {
	character Character
}

func (f *Finery) Decorate(ch Character) {
	f.character = ch
}

func (f *Finery) Show() {
	f.character.Show()
}
