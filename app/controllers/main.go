package controllers
import (
	//"github.com/astaxie/beego"
	"github.com/johnstevin/goblog/app/models"
	"strings"
	"fmt"
)

type MainController struct {
	BaseController
}

func (this *MainController) Login() {
	if this.isPost() {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")
		fmt.Printf("username:%s,password:%s,remember:%s--|",username,password,remember)
		if username != "" && password != "" {
			user, err := models.GetUserByName(username)
			fmt.Printf("user:%s,err:%s--|",user,err)
//			errorMsg := ""
//			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
//				errorMsg = "帐号或密码错误"
//			} else if user.Status == -1 {
//				errorMsg = "该帐号已禁用"
//			} else {
//				if remember == "1" {
//					fmt.Printf("remember = 1")
//				} else {
//					fmt.Printf("remember != 1")
//				}
//				this.redirect(beego.URLFor("UserController.Index"))
//			}
			
		}
		
	}
	this.TplName = "main/login.html"
}