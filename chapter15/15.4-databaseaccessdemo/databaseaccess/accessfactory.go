package databaseaccess

type AccessFactory struct {
}

func (s *AccessFactory) CreateUser() UserInterface {
	return &AccessUser{}
}

func (s *AccessFactory) CreateDepartment() DepartmentInterface {
	return &AccessDepartment{}
}
