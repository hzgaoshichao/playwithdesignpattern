package basketballtranslator

type PlayerInterface interface {
	Attack()
	Defense()
}

type player struct {
	name string
}
