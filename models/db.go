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
package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hunterhug/GoWeb/models/admin"
	_ "github.com/hunterhug/GoWeb/models/blog"
	"os"
	"time"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

func Createtb() {
	beego.Trace("data init start")
	admin.InitData()
	beego.Trace("data init end")
}

func Syncdb(force bool) {
	beego.Trace("db, sync db start")

	Createdb(force)
	Connect()
	Createconfig()
	Createtb()

	beego.Trace("sync db end, please reopen app again")
}

func Updaterbac() {
	TRUNCATETable([]string{beego.AppConfig.String("rbac_group_table"), beego.AppConfig.String("rbac_node_table")})
	Connect()
	admin.InsertGroup()
	admin.InsertNodes()
}

func Createconfig() {
	name := "default" // database alias name
	force := true     // drop table force
	verbose := true   // print log
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("database config set to force error:" + err.Error())
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
		orm.DefaultTimeLoc = time.UTC
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

//创建数据库
func Createdb(force bool) {
	beego.Trace("create database start")
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	// db_path := beego.AppConfig.String("db_path")
	// db_sslmode := beego.AppConfig.String("db_sslmode")

	var dns string
	var sqlstring, sql1string string

	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
		sqlstring = fmt.Sprintf("CREATE DATABASE if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
		sql1string = fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", db_name)
		if force {
			fmt.Println(sql1string)
		}
		fmt.Println(sqlstring)
		break
	// case "postgres":
	// 	dns = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_host, db_user, db_pass, db_port, db_sslmode)
	// 	sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
	// 	break
	// case "sqlite3":
	// 	if db_path == "" {
	// 		db_path = "./"
	// 	}
	// 	dns = fmt.Sprintf("%s%s.db", db_path, db_name)
	// 	os.Remove(dns)
	// 	sqlstring = "create table init (n varchar(32));drop table init;"
	// 	break
	default:
		beego.Critical("db driver not support:", db_type)
		return
	}
	db, err := sql.Open(db_type, dns)
	if err != nil {
		panic(err.Error())
	}
	if force {
		_, err = db.Exec(sql1string)
	}
	_, err1 := db.Exec(sqlstring)
	if err != nil || err1 != nil {
		beego.Error("db exec error：", err, err1)
		panic(err.Error())
	} else {
		beego.Trace("database ", db_name, " created")
	}
	defer db.Close()
	beego.Trace("create database end")
}

func TRUNCATETable(table []string) {
	beego.Trace("delete tables start")
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	var dns string
	var sqlstring string
	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		// case "postgres":
		// 	dns = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_host, db_user, db_pass, db_port, db_sslmode)
		// 	sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
		// 	break
		// case "sqlite3":
		// 	if db_path == "" {
		// 		db_path = "./"
		// 	}
		// 	dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		// 	os.Remove(dns)
		// 	sqlstring = "create table init (n varchar(32));drop table init;"
		// 	break
	default:
		beego.Critical("db driver not support:", db_type)
		return
	}
	db, err := sql.Open(db_type, dns)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	for _, i := range table {
		beego.Trace("table deleting：" + i)
		sqlstring = fmt.Sprintf("TRUNCATE TABLE `%s`", i)
		_, err = db.Exec(sqlstring)
		if err != nil {
			beego.Error("table delete error：" + err.Error())
			panic(err.Error())
		} else {
			beego.Trace("table delete success：" + i)
		}
	}
	beego.Trace("delete table end")
}
