package databaseaccess

type FactoryInterface interface {
	CreateUser() User
}
