package databaseaccess

type SqlserverFactory struct {
}

func (s *SqlserverFactory) CreateUser() UserInterface {
	return &SqlserverUser{}
}

func (s *SqlserverFactory) CreateDepartment() DepartmentInterface {
	return &SqlserverDepartment{}
}
