package gameprogress

import "fmt"

type Gamerole struct {
	vitality int
	attack   int
	defense  int
}

func (g *Gamerole) getVitality() int {
	return g.vitality
}

func (g *Gamerole) setVitality(value int) {
	g.vitality = value
}

func (g *Gamerole) getAttack() int {
	return g.attack
}

func (g *Gamerole) setAttack(value int) {
	g.attack = value
}

func (g *Gamerole) getDefense() int {
	return g.defense
}

func (g *Gamerole) setDefense(value int) {
	g.defense = value
}

func (g *Gamerole) DisplayState() {
	fmt.Printf("角色当前状态: \n")
	fmt.Printf("体力: %v \n", g.vitality)
	fmt.Printf("攻击力: %v \n", g.attack)
	fmt.Printf("防御力: %v \n", g.defense)
}

func (g *Gamerole) GetInitState() {
	g.vitality = 100
	g.attack = 100
	g.defense = 100
}

func (g *Gamerole) Fight() {
	g.vitality = 0
	g.attack = 0
	g.defense = 0
}

func (g *Gamerole) SaveState() *roleStateMemento {

	return NewRoleStateMemento(g.vitality, g.attack, g.defense)

}

func (g *Gamerole) RecoverState(rsm *roleStateMemento) {
	g.vitality = rsm.vitality
	g.attack = rsm.attack
	g.defense = rsm.defense
}
