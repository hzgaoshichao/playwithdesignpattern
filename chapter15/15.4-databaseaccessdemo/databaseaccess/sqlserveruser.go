package databaseaccess

import "fmt"

type SqlserverUser struct {
}

func (s *SqlserverUser) Insert(user User) {
	fmt.Printf("在 Sql Server 中给 User 增加了一条记录 \n")
}
func (s *SqlserverUser) GetUser(id int) *User {
	fmt.Printf("在 Sql Server 中根据用户 id 得到 User 表的一条记录 \n")
	return nil
}
