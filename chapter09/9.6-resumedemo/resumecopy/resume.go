package resumecopy

import "fmt"

type Resume struct {
	name string
	sex  string
	age  string
	work *workExperience
}

func NewResume(name string) Resume {
	return Resume{
		name: name,
		work: new(workExperience),
	}

}

func (r *Resume) SetPersonalInfo(sex string, age string) {
	r.sex = sex
	r.age = age
}

func (r *Resume) SetWorkExperience(timeArea string, company string) {
	r.work.SetTimeArea(timeArea)
	r.work.SetCompany(company)

}

func (r *Resume) Display() {
	fmt.Printf("name: %v, sex: %v, age: %v \n", r.name, r.sex, r.age)
	fmt.Printf("工作经历 %v %v \n", r.work.timeArea, r.work.company)
}

func (r *Resume) Clone() cloneable {
	newResume := *r
	newResume.work = r.work.Clone().(*workExperience)
	return &newResume
}
