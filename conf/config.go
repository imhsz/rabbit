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
package conf

import (
	"github.com/astaxie/beego"
	"strconv"
)

type FlagConfig struct {
	User        *string
	DbInit      *bool
	DbInitForce *bool
	Rbac        *bool
}

var (
	AuthType     int
	AuthGateWay  string
	AuthAdmin    string
	Cookie7      bool
	Version      string
	HomeTemplate string
)

func InitConfig() {
	Version = beego.AppConfig.DefaultString("version", "HG V4")
	AuthType, _ = strconv.Atoi(beego.AppConfig.String("user_auth_type"))
	AuthGateWay = beego.AppConfig.DefaultString("rbac_auth_gateway", "/public/login")
	Cookie7, _ = beego.AppConfig.Bool("cookie7")
	AuthAdmin = beego.AppConfig.DefaultString("rbac_admin_user", "admin")
	HomeTemplate = beego.AppConfig.DefaultString("home_template", "default")
}
