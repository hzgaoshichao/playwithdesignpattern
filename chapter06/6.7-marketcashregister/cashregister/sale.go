package cashregister

type Sale interface {
	acceptCash(price float64, num int) float64
}
