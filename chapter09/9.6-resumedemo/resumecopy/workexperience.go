package resumecopy

type workExperience struct {
	timeArea string
	company  string
}

func (w *workExperience) SetTimeArea(value string) {
	w.timeArea = value
}
func (w *workExperience) SetCompany(value string) {
	w.company = value
}

func (w *workExperience) Clone() cloneable {
	newWorkExperience := *w
	return &newWorkExperience
}
