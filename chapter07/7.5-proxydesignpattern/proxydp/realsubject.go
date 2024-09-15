package proxydp

import "fmt"

type RealSubject struct {
}

func (p *RealSubject) Request() {
	fmt.Println("真实的请求。")
}
