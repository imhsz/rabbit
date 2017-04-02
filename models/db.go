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

var o orm.Ormer

// 建数据库表
func Createtb() {
	beego.Trace("开始数据填充")
	admin.InitData()
	beego.Trace("数据填充结束")
}

// 同步数据库
func Syncdb() {
	beego.Trace("-s,数据库同步开始")
	Createdb()
	Connect()
	o = orm.NewOrm()
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("同步数据库错误：", err)
	}
	Createtb()
	beego.Trace("数据库同步完毕，请重新开启应用")

}

// 数据库连接
func Connect() {
	beego.Trace("数据库开始连接注册")
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
	// db_path := beego.AppConfig.String("db_path")
	// db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.DefaultTimeLoc = time.UTC
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	// case "postgres":
	// 	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 	dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
	// case "sqlite3":
	// 	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	// 	if db_path == "" {
	// 		db_path = "./"
	// 	}
	// 	dns = fmt.Sprintf("%s%s.db", db_path, db_name)
	// 	break
	default:
		beego.Critical("数据库驱动暂不支持：", db_type)
	}
	err := orm.RegisterDataBase("default", db_type, dns)
	if err != nil {
		beego.Error("注册数据库失败：" + err.Error())
		panic(err.Error())
	}

	if beego.AppConfig.String("dblog") == "open" {
		beego.Trace("应用开发者模式，数据库操作进行调试，记录进db.log")
		orm.Debug = true
		w, e := os.OpenFile("log/db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if e != nil {
			beego.Error(e.Error())
		}
		orm.DebugLog = orm.NewLog(w)
	}

}

//创建数据库
func Createdb() {
	beego.Trace("开始建库")
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	// db_path := beego.AppConfig.String("db_path")
	// db_sslmode := beego.AppConfig.String("db_sslmode")

	var dns string
	var sqlstring string
	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
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
		beego.Critical("数据库驱动暂不支持：", db_type)
	}
	db, err := sql.Open(db_type, dns)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(sqlstring)
	if err != nil {
		beego.Error("数据库操作执行失败：" + err.Error())
		panic(err.Error())
	} else {
		beego.Trace("Database ", db_name, " created")
	}
	defer db.Close()
	beego.Trace("建库结束")
}
