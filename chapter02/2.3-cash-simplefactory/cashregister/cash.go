package cashregister

import "math"

type Cash interface {
	// 收取费用的接口方法,参数为单价和数量
	AcceptCash(price float64, num int) float64
}

type CashNormal struct {
}

func NewCashNormal() *CashNormal {
	return &CashNormal{}
}

// 正常收费
func (cn *CashNormal) AcceptCash(price float64, num int) float64 {
	return price * float64(num)
}

// 打折收费
type CashRebate struct {
	moneyRebate float64
}

func NewCashRebate(moneyRebate float64) *CashRebate {
	return &CashRebate{
		moneyRebate: moneyRebate,
	}
}

func (cr *CashRebate) AcceptCash(price float64, num int) float64 {
	return price * float64(num) * cr.moneyRebate
}

// 满返收费,比如满300,减去100
type CashReturn struct {
	//返利条件, 比如满300,减去100
	moneyCondition int
	moneyReturn    int
}

func NewCashReturn(moneyCondition int, moneyReturn int) *CashReturn {
	return &CashReturn{
		moneyCondition: moneyCondition,
		moneyReturn:    moneyReturn,
	}
}

func (cr *CashReturn) AcceptCash(price float64, num int) float64 {
	result := price * float64(num)
	if cr.moneyCondition > 0 && result >= float64(cr.moneyCondition) {
		result = result - math.Floor(float64(result)/float64(cr.moneyCondition))*float64(cr.moneyReturn)
	}
	return result
}

func CreateCashAccept(cashType int) Cash {
	var cash Cash
	switch cashType {
	case 1:
		cash = NewCashNormal() // 正常收费
	case 2:
		cash = NewCashRebate(0.8) // 打八折
	case 3:
		cash = NewCashRebate(0.7) // 打七折
	case 4:
		cash = NewCashReturn(300, 100) // 满 300 返 100
	}
	return cash
}
