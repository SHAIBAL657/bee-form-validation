package routers

import (
	"bee-form-validation/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ObjectController{})
}
