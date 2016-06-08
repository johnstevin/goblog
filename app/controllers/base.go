package controllers
import (
	"github.com/astaxie/beego"
	"strings"
	//"fmt"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	//fmt.Printf("%s",controllerName)
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName     = strings.ToLower(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
}

func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}
	this.Layout = "layout/layout.html"
	this.TplName = tplname
}