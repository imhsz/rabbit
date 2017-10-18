/*
   Created by jinhan on 17-10-18.
   Tip:
   Update:
*/
package home

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"testing"
	"time"
)

func init() {
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	Connect()
}

// in hear
func TestGetNav(t *testing.T) {
	a := GetNav(0, 0)
	b, _ := json.Marshal(a)
	fmt.Printf("%v", string(b))
}

func Connect() {
	beego.Trace("database start to connect")
	var dns string
	db_type := beego.AppConfig.String("db_type")
	beego.Trace("db type " + db_type)
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
		beego.Trace("dns " + dns)
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
