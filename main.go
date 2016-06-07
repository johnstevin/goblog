package main
import (
	"github.com/astaxie/beego"
	"github.com/johnstevin/goblog/app/models"
	"net/http"
	"html/template"
)

const VERSION = "1.0.0"

func main() {
	beego.ErrorHandler("404", page_not_found)
	beego.ErrorHandler("dbError", dbError)
	
	models.Init()
	beego.AppConfig.Set("version", VERSION)
	beego.Run()
}

func page_not_found(rw http.ResponseWriter, r *http.Request){
    t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
    data := make(map[string]interface{})
    data["content"] = "page not found"
    t.Execute(rw, data)
}

func dbError(rw http.ResponseWriter, r *http.Request){
    t,_:= template.New("dberror.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/dberror.html")
    data :=make(map[string]interface{})
    data["content"] = "database is now down"
    t.Execute(rw, data)
}