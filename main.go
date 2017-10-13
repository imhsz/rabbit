/*
	Copyright 2017 by GoWeb author: gdccmcm14@live.com.
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License
*/

// Main Web Entrance
package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/hunterhug/GoWeb/conf"
	"github.com/hunterhug/GoWeb/controllers"
	"github.com/hunterhug/GoWeb/lib"
	"github.com/hunterhug/GoWeb/models"
	"github.com/hunterhug/GoWeb/routers"
	"mime"
	"strings"
)

// 国际化语言数组
var langTypes []string
var home *string

// 加载、初始化国际化
func init() {
	flags := conf.FlagConfig{}
	config := flag.String("config", "", "config file position if empty use default")
	flags.User = flag.String("user", "", "user")
	flags.DbInit = flag.Bool("db", false, "init db")
	flags.DbInitForce = flag.Bool("f", false, "force init db first drop db then rebuild it")
	flags.Rbac = flag.Bool("rbac", false, "rebuild rbac database tables")
	home = flag.String("home", "", "home template")

	flag.Parse()

	if *config != "" {
		beego.Trace("use diy config")
		err := beego.LoadAppConfig("ini", *config)
		if err != nil {
			beego.Trace(err.Error())
		} else {
			beego.Trace("Use config:" + *config)
		}
	}

	// just add some ini in conf such locale_zh-CN.ini and edit app.conf
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	for _, lang := range langTypes {
		beego.Trace("Load language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Load language error:", err)
			return
		}
	}

	// 添加映射
	beego.Trace("add i18n function map")
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Trace("add stringsToJson function  map")
	beego.AddFuncMap("stringsToJson", lib.StringsToJson)
	mime.AddExtensionType(".css", "text/css") // some not important

	// 模型初始化
	beego.Trace("model run")
	models.Run(flags)

	beego.Trace("router run")
	routers.Run()

	beego.Trace("start open error template")
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	beego.Trace("Start Listen ...")
	conf.InitConfig()

	if *home != "" {
		beego.Trace("Home template is " + *home)
		beego.AppConfig.Set(beego.BConfig.RunMode+"::"+"home_template", *home)
	}
	beego.Run()
}
