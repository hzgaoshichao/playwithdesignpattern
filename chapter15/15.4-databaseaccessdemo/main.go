package main

import dba "databaseaccessdemo/databaseaccess"

func main() {
	user := dba.User{}
	factory := dba.AccessFactory{} // SqlserverFactory{}
	iu := factory.CreateUser()
	iu.Insert(user)
	iu.GetUser(1)

	dp := dba.Department{}
	idep := factory.CreateDepartment()
	idep.Insert(dp)
	idep.GetDepartment(1)
}
