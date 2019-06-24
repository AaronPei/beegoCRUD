package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"] = append(beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"] = append(beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"] = append(beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"] = append(beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"] = append(beego.GlobalControllerRouter["learn/beegoTest/nodeApi/controllers:NodesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
