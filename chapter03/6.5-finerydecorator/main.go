package main

import (
	pf "finerydecorator/personfinery"
	"fmt"
)

func main() {
	xc := pf.NewPerson("小菜")

	fmt.Println("第一种装扮......")

	pqx := new(pf.Sneakers)
	pqx.Decorate(&xc)

	kk := pf.BigTrouser{}
	kk.Decorate(pqx)

	dtx := pf.Tshirts{}
	dtx.Decorate(&kk)

	dtx.Show()

	fmt.Println("第二种装扮......")
	px := new(pf.LeatherShoes)
	px.Decorate(&xc)

	ld := new(pf.Tie)
	ld.Decorate(px)

	xz := new(pf.Suit)
	xz.Decorate(ld)

	xz.Show()

	fmt.Println("第三种装扮......")
	pqx2 := new(pf.Sneakers)
	pqx2.Decorate(&xc)

	px2 := new(pf.LeatherShoes)
	px2.Decorate(pqx2)

	kk2 := new(pf.BigTrouser)
	kk2.Decorate(px2)

	ld2 := new(pf.Tie)
	ld2.Decorate(kk2)

	cm2 := new(pf.Strawhat)
	cm2.Decorate(ld2)

	cm2.Show()

}
