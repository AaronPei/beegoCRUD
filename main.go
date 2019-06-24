package main

import (
	"learn/beegoTest/nodeApi/credentials"
	_ "learn/beegoTest/nodeApi/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func main() {
	credentials.Init_Credentials()
	beego.Run()
}
