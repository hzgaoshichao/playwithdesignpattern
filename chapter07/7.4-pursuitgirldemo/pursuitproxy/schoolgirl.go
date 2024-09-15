package pursuitproxy

type SchoolGirl struct {
	name string
}

func (sg *SchoolGirl) SetName(name string) {
	sg.name = name
}
