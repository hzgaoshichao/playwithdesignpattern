package main

import g "gameprogressmemento/gameprogress"

func main() {
	role := g.Gamerole{}
	role.GetInitState()
	role.DisplayState()

	stateAdmin := g.RoleStateMementoCaretaker{}
	stateAdmin.SetRoleStateMemento(role.SaveState())

	role.Fight()
	role.DisplayState()

	role.RecoverState(stateAdmin.GetRoleStateMemento())
	role.DisplayState()

}
