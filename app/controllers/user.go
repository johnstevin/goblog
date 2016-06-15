package controllers
import (
	//"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

func (this *UserController) Index() {
	this.Data["pageTitle"] = "个人主页"
	this.Data["user"] = "stevin"
	this.display()
}

func (this *UserController) List() {
	this.Data["pageTitle"] = "好友列表"
	this.display()
}

func (this *UserController) Profile() {
	
}