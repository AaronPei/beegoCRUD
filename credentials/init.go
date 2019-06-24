package credentials

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init_Credentials() {
	orm.RegisterDriver(beego.AppConfig.String("ConnectionType"), orm.DRMySQL)
	orm.RegisterDataBase("default",
		beego.AppConfig.String("ConnectionType"),
		beego.AppConfig.String("Username")+
			":"+beego.AppConfig.String("Password")+
			"@tcp("+beego.AppConfig.String("Host")+
			":"+beego.AppConfig.String("Port")+
			")/"+beego.AppConfig.String("DBName")+"?charset=utf8")
	fmt.Printf(beego.AppConfig.String("Username") +
		":" + beego.AppConfig.String("Password") +
		"@tcp(" + beego.AppConfig.String("Host") +
		":" + beego.AppConfig.String("Port") +
		")/" + beego.AppConfig.String("DBName") + "?charset=utf8\n")
	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
