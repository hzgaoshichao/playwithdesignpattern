package databaseaccess

import "fmt"

type AccessUser struct {
}

func (s *AccessUser) Insert(user User) {
	fmt.Printf("在 Access 中给 User 增加了一条记录 \n")
}
func (s *AccessUser) GetUser(id int) *User {
	fmt.Printf("在 Access 中根据用户 id 得到 User 表的一条记录 \n")
	return nil
}
