package iteratorpattern

type IteratorInterface interface {
	First() any
	Next() any
	IsDone() bool
	CurrentItem() any
}
