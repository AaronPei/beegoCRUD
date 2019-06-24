package routers

import (
	"learn/beegoTest/nodeApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/default", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/nodes",
			beego.NSInclude(
				&controllers.NodesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
