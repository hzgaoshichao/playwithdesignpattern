package adapterpattern

type adapter struct {
	adp *adaptee
}

func (a *adapter) Request() {
	a.adp.specificRequest()
}

func NewAdapter() *adapter {
	return &adapter{
		adp: &adaptee{},
	}
}
