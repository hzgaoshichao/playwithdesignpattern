package cashregister

type factory interface {
	createSalesMode() Sale
}
