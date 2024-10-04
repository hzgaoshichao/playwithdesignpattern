package main

import co "observerdemo/companyobserver"

func main() {

	boss := co.NewBoss("胡汉三")
	employee1 := co.NewStockObserver("魏关姹", boss)
	employee2 := co.NewStockObserver("易管查", boss)
	employeelqm := co.NewNBAObserver("兰秋冥", boss)
	emphardworker := co.NewNBAObserver("打工人", boss)

	boss.Attach(employee1)
	boss.Attach(employee2)
	boss.Attach(emphardworker)
	boss.Attach(employeelqm)
	boss.Detach(emphardworker)
	boss.SetAction("我胡汉三回来了")
	boss.NotifyEmployee()

	secretary01 := co.NewSecretary("秘书A")
	employee3 := co.NewStockObserver("魏关姹", secretary01)
	secretary01.Attach(employee3)
	secretary01.SetAction("老板回来了")
	secretary01.NotifyEmployee()

}
