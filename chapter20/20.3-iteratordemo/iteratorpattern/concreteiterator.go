package iteratorpattern

type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	current   int
}

func NewConcreteIterator(agg *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{
		aggregate: agg,
		current:   0,
	}
}

func (c *ConcreteIterator) First() any {
	return c.aggregate.getCurrentItem(0)
}

func (c *ConcreteIterator) Next() any {
	var ret any
	c.current++
	if c.current < c.aggregate.getCount() {

		ret = c.aggregate.getCurrentItem(c.current)
	}
	return ret
}

func (c *ConcreteIterator) IsDone() bool {
	if c.current >= c.aggregate.getCount() {
		return true
	} else {
		return false

	}
}

func (c *ConcreteIterator) CurrentItem() any {
	return c.aggregate.getCurrentItem(c.current)
}
