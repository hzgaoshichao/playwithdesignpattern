package cashregister

type cashRebateReturnFactory struct {
	moneyRebate    float64
	moneyCondition int
	moneyReturn    int
}

func NewCashRebateReturnFactory(moneyRebate float64, moneyCondition int, moneyReturn int) *cashRebateReturnFactory {
	return &cashRebateReturnFactory{
		moneyRebate:    moneyRebate,
		moneyCondition: moneyCondition,
		moneyReturn:    moneyReturn,
	}
}

func (crr *cashRebateReturnFactory) createSalesMode() Sale {
	//先打折, 再满返
	cn := NewCashNormal()
	cret := newCashReturn(crr.moneyCondition, crr.moneyReturn)
	creb := newCashRebate(crr.moneyRebate)

	cret.decorate(cn)
	creb.decorate(cret)

	return creb

}
