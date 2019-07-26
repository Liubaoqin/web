package routers

import (
	"web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("index", &controllers.IndexController{}, "*:Index")
	beego.Router("about", &controllers.IndexController{}, "*:About")
	beego.Router("message", &controllers.IndexController{}, "*:Message")
	beego.Router("details", &controllers.IndexController{}, "*:Details")
	beego.Router("comment", &controllers.IndexController{}, "*:Comment")
	beego.Router("error", &controllers.IndexController{}, "*:Error")

	beego.Router("cms", &controllers.CmsController{})
}
