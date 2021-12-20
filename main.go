package main

import (
	_ "bee-form-validation/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/"] = "views/index.tpl"
	}
	beego.Run()
}
