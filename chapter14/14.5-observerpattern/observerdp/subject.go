package observerdp

import "reflect"

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	NotifyObserver()
	GetSubjectState() string
	SetSubjectState(ss string)
}

type SubjectCommon struct {
	list         []Observer
	subjectState string
}

func (s *SubjectCommon) GetSubjectState() string {
	return s.subjectState
}

func (s *SubjectCommon) SetSubjectState(ss string) {
	s.subjectState = ss
}

func (s *SubjectCommon) Attach(observer Observer) {
	s.list = append(s.list, observer)
}

func (s *SubjectCommon) Detach(observer Observer) {
	for i, ob := range s.list {
		if reflect.DeepEqual(observer, ob) {
			s.list = append(s.list[:i], s.list[i+1:]...)
		}
	}
}

func (s *SubjectCommon) NotifyObserver() {
	for _, ob := range s.list {
		ob.update()
	}
}

type ConcreteSubject struct {
	SubjectCommon
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		SubjectCommon: SubjectCommon{
			list: make([]Observer, 0),
		},
	}
}
