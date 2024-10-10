package cms

import "fmt"

type financeDepartment struct {
	name string
}

func NewFinanceDepartment(name string) *financeDepartment {
	return &financeDepartment{name: name}
}

func (f *financeDepartment) Add(com companyInterface) {

}

func (f *financeDepartment) Remove(com companyInterface) {

}

func (f *financeDepartment) Display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("%v \n", f.name)

}

func (f *financeDepartment) LineOfDuty() {

	fmt.Printf("%v 公司财务收支管理 \n", f.name)
}
