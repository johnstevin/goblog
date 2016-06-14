package controllers
import (
	"github.com/astaxie/beego"
	"github.com/johnstevin/goblog/app/models"
	"github.com/johnstevin/goblog/app/libs"
	"strings"
	"time"
	"strconv"
)

type MainController struct {
	BaseController
}

func (this *MainController) Login() {
	beego.ReadFromRequest(&this.Controller)
	
	if this.isPost() {
		flash := beego.NewFlash()

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")
		if username != "" && password != "" {
			errorMsg := ""
			user, err := models.GetUserByName(username)
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "账号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "账号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				err := models.UserUpdate(user,"LastIp","LastLogin")
				if err != nil {
					errorMsg = "登录更新失败"
				} else {
					authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
					if remember == "1" {
						this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
					} else {
						this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey)
					}
					this.redirect(beego.URLFor("UserController.Index"))
				}
				
			}

			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("MainController.Login"))
		}
		
	}
	this.TplName = "main/login.html"
}