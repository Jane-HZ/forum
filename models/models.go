package models

import (
	"errors"
	"fmt"
	_"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
    "log"
)

type User struct {
	Id       int
	Name     string `orm:"size(50)"`
	Password string `orm:"size(50)"`
}


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //设置驱动
	orm.RegisterDataBase("default", "mysql", "root:ac1998729@tcp(127.0.0.1:3306)/my_blog?charset=utf8", 30)//连接数据库
	orm.RegisterModel(new(User)) //注册模型
}

func (user User) GetInformation() (int,error) {
	orm.Debug = true
	o := orm.NewOrm()
	if err := o.Read(&user,"name","password"); err != nil  {
		return 0, errors.New("用户名或密码错误")
	} else {
		return user.Id,nil
	}
}

func (user User) SaveInformation() (int, error) {
	o := orm.NewOrm()
	if created, id, err := o.ReadOrCreate(&user,"name"); err == nil {
		if created {
			return int(id) , nil
		} else {
			user.Id = int(id)
			if num, err := o.Update(&user); err == nil && num > 0 {
				return int(id) , nil
			}
		}
	}
	return 0 , fmt.Errorf("save fail")
}

func (user *User)GetUserId() error {
	o := orm.NewOrm()
	if user.Name == "" {
		return errors.New("用户未登录，请先登录")
	}
	if err := o.Read(user, "name"); err != nil {
		log.Printf("read user %v error,error info is %v \n", user, err)
		return errors.New("用户不存在")
	}
	return nil
}

