package routers

import "github.com/astaxie/beego"

 import (
// 	"github.com/astaxie/beego"
// 	"github.com/hunterhug/GoWeb/controllers/admin"
 	"github.com/hunterhug/GoWeb/controllers/home"
 )

func homerouter() {

 	//前台路由
 	beego.Router("/", &home.MainController{}, "*:Index")
 	beego.Router("/:id/", &home.MainController{}, "*:Category")
	beego.Router("/:cid/:id/", &home.MainController{}, "*:Paper")
 }
