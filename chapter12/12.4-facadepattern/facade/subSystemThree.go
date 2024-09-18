package facade

import "fmt"

type subSystemThree struct {
}

func (s subSystemThree) methodThree() {
	fmt.Printf("子系统方法三 \n")
}
