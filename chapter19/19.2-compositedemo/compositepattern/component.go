package compositepattern

type componentInterface interface {
	Add(comp componentInterface)
	Remove(com componentInterface)
	Display(depth int)
}

type component struct {
	name string
}
