package cms

type companyInterface interface {
	Add(com companyInterface)
	Remove(com companyInterface)
	Display(depth int)
	LineOfDuty()
}
