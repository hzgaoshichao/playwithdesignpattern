package gameprogress

type RoleStateMementoCaretaker struct {
	rsm *roleStateMemento
}

func (r *RoleStateMementoCaretaker) GetRoleStateMemento() *roleStateMemento {
	return r.rsm
}
func (r *RoleStateMementoCaretaker) SetRoleStateMemento(value *roleStateMemento) {
	r.rsm = value
}
