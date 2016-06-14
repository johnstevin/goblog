package models
import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id           int
	UserName     string
	Email		 string
	Password     string
	Salt  		 string
	LastLogin    int64
	LastIp       string
	Status       int
}

func (u *User) TableName() string {
	return TableName("user")
}

func GetUserByName(userName string) (*User, error) {
	u := new(User)
	err := orm.NewOrm().QueryTable(u.TableName()).Filter("user_name", userName).One(u)
	if err != nil {
		return nil, err
	}
	return u,nil
}

func UserUpdate(user *User, fields ...string) error {
	_, err := orm.NewOrm().Update(user, fields...)
	return err
}