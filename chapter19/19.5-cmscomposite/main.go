package main

import (
	"fmt"

	cms "cmscomposite/cms"
)

func main() {
	root := cms.NewConcreteCompany("北京总公司")
	root.Add(cms.NewFinanceDepartment("总公司财务部"))
	root.Add(cms.NewHRDepartment("总公司人力资源"))

	comp := cms.NewConcreteCompany("上海华东分公司")
	comp.Add(cms.NewHRDepartment("华东分公司人力资源部"))
	comp.Add(cms.NewFinanceDepartment("华东分公司财务部"))
	root.Add(comp)

	comp2 := cms.NewConcreteCompany("南京办事处")
	comp2.Add(cms.NewHRDepartment("南京办事处人力资源部"))
	comp2.Add(cms.NewFinanceDepartment("南京办事处财务部"))
	comp.Add(comp2)

	comp3 := cms.NewConcreteCompany("杭州办事处")
	comp3.Add(cms.NewHRDepartment("杭州办事处人力资源部"))
	comp3.Add(cms.NewFinanceDepartment("杭州办事处财务部"))
	comp.Add(comp3)

	fmt.Printf("结构图: \n")
	root.Display(1)
	fmt.Printf("职责: \n")
	root.LineOfDuty()
}
