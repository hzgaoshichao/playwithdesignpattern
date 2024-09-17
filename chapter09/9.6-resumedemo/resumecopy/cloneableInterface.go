package resumecopy

type cloneable interface {
	Clone() cloneable
}
