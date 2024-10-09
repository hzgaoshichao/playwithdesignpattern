package basketballtranslator

type translator struct {
	f *foreignCenter
}

func NewTranslator(name string) *translator {
	return &translator{
		f: NewForeignCenter(name),
	}
}

func (t *translator) Attack() {
	t.f.foreignCenterAttack()
}

func (t *translator) Defense() {
	t.f.foreignCenterDefense()
}
