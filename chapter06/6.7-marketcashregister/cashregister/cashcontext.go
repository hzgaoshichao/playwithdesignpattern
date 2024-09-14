package cashregister

type CashContext struct {
	sale Sale
}

func NewCashContext(cashType int) CashContext {
	var cs Sale
	switch cashType {
	case 1:
		cs = NewCashNormal()
	case 2:
		cs = newCashRebate(0.8)
	case 3:
		cs = newCashRebate(0.7)
	case 4:
		cs = newCashReturn(300, 100)
	case 5:
		//先打8折, 再满300返100
		cn := NewCashNormal()
		cr1 := newCashReturn(300, 100)
		cr2 := newCashRebate(0.8)
		cr1.decorate(cn)
		cr2.decorate(cr1)
		cs = cr2
	case 6:
		//先满200返50, 再打7折
		cn2 := NewCashNormal()
		cr3 := newCashRebate(0.7)
		cr4 := newCashReturn(200, 50)
		cr3.decorate(cn2)
		cr4.decorate(cr3)
		cs = cr4
	}

	return CashContext{
		sale: cs,
	}
}

func (cc CashContext) GetResult(price float64, num int) float64 {
	return cc.sale.acceptCash(price, num)
}
