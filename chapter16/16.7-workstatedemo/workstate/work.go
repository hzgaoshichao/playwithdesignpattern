package workstate

type Work struct {
	current      StateInterface
	hour         int
	workFinished bool
}

func NewWork() *Work {
	return &Work{
		current: &ForenoonState{},
	}
}

func (w *Work) SetState(s StateInterface) {
	w.current = s
}

func (w *Work) WriteProgram() {
	w.current.writeProgram(w)
}

func (w *Work) GetHour() int {
	return w.hour
}

func (w *Work) SetHour(h int) {
	w.hour = h
}

func (w *Work) GetWorkFinished() bool {
	return w.workFinished
}

func (w *Work) SetWorkFinished(value bool) {
	w.workFinished = value
}
