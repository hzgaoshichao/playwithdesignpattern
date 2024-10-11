package iteratorpattern

type AggregateInterface interface {
	CreateIterator() IteratorInterface
}
