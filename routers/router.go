package routers

import (
	"mvc-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/users", &controllers.UserController{})
    beego.Router("/users/list", &controllers.UserController{}, "get:List")
}
