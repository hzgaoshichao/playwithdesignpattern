package proxydp

type Proxy struct {
	subject Subject
}

func NewProxy() *Proxy {
	return &Proxy{subject: &RealSubject{}}
}

func (p *Proxy) Request() {
	p.subject.Request()
}
