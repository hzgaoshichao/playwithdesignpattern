package cashregister

type CashSuper struct {
	sale Sale
}

func (cs *CashSuper) decorate(sale Sale) {
	cs.sale = sale

}

func (cs *CashSuper) acceptCash(price float64, num int) float64 {
	var result float64
	if cs.sale != nil {
		result = cs.sale.acceptCash(price, num)
	}
	return result
}
