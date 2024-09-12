package cashregister

type Cash interface {
	// 收取费用的接口方法,参数为单价和数量
	acceptCash(price float64, num int) float64
}
