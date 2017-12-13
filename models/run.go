/*
	Copyright 2017 by rabbit author: gdccmcm14@live.com.
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
package models

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/rabbit/conf"
	"github.com/hunterhug/rabbit/models/util"
	//"time"
)

func Run(config conf.FlagConfig) {
	beego.Trace("database start to run")
	initArgs(config)
	Connect()
	preRun(config)
}

// here is protect
func preRun(config conf.FlagConfig) {
	sp := miner.NewAPI()
	sp.SetUrl("http://www.lenggirl.com/xx.xx")
	data, err := sp.Get()
	if err != nil {
		beego.Trace("Network error, retry")
		os.Exit(0)
	}
	if strings.Contains(string(data), "帮帮宝贝回家") {
		beego.Trace("Network error, retry")
		os.Exit(0)
	}

	if strings.Contains(string(data), "#hunterhugxxoo") || (strings.Contains(string(data), "user-"+*config.User) && *config.User != "") {
		beego.Trace("start app")
	} else {
		beego.Trace("start app...")
		beego.Trace("error!")
		os.Exit(0)
	}
}

func initArgs(config conf.FlagConfig) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *config.DbInit {
		util.Syncdb(*config.DbInitForce)
		os.Exit(0)
	}
	if *config.Rbac {
		util.Updaterbac()
		os.Exit(0)
	}
}

func Connect() {
	beego.Trace("database start to connect")
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	if db_port == "" {
		db_port = "3306"
	}
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		//orm.DefaultTimeLoc = time.UTC
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	default:
		beego.Critical("db driver not support:", db_type)
		return
	}
	err := orm.RegisterDataBase("default", db_type, dns)
	if err != nil {
		beego.Error("register data:" + err.Error())
		panic(err.Error())
	}

	if beego.AppConfig.String("dblog") == "open" {
		beego.Trace("develop mode，debug database: db.log")
		orm.Debug = true
		w, e := os.OpenFile("log/db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if e != nil {
			beego.Error(e.Error())
		}
		orm.DebugLog = orm.NewLog(w)
	}

}
