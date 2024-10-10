package cms

import (
	"fmt"
)

type hrDepartment struct {
	name string
}

func NewHRDepartment(name string) *hrDepartment {
	return &hrDepartment{
		name: name,
	}
}

func (h *hrDepartment) Add(com companyInterface) {

}

func (h *hrDepartment) Remove(com companyInterface) {

}

func (h *hrDepartment) Display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%v \n", h.name)

}

func (h *hrDepartment) LineOfDuty() {

	fmt.Printf("%v 员工招聘培训管理 \n", h.name)
}
