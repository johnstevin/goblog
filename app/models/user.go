package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type User struct {
	Id           int
	UserName     string
	Email		 string
	Password     string
	Salt  		 string
	LastLogin    int
	lastIp       string
	Status       int
}

func (u *User) TableName() string {
	return TableName("user")
}

func GetUserByName(userName string) (*User, error) {
	u := new(User)
	fmt.Printf("%s--|",u)
	tname := u.TableName()
	fmt.Printf("%s--|",tname)
	err := orm.NewOrm().QueryTable(u.TableName()).Filter("user_name", userName).One(u)
	if err != nil {
		return nil, err
	}
	return u,nil
//	err := orm.NewOrm().QueryTable(TableName("user")).Filter("user_name", userName).One(u)
//	if err != nil {
//		return nil, err
//	}
//	return u, nil
}