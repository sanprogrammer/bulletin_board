package routers

import (
	"bulletin_board/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/animals/:id([0-9]+", &controllers.MainController{}, "Get:Show")
}
