package rbac

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// 基础控制器
type baseController struct {
	beego.Controller
	i18n.Locale
}

type Tree struct {
	Id         int64      `json:"id"`
	GroupId    int64      `json:"-"`
	Text       string     `json:"text"`
	IconCls    string     `json:"iconCls"`
	Checked    string     `json:"checked"`
	State      string     `json:"state"`
	Children   []Tree     `json:"children"`
	Attributes Attributes `json:"attributes"`
}
type Attributes struct {
	Url   string `json:"url"`
	Price int64  `json:"price"`
}

// 通过浏览器获取语言
func (this *baseController) Prepare() {
	this.Lang = ""

	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5]
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}

	if len(this.Lang) == 0 {
		this.Lang = "zh-CN"
	}

	this.Data["Lang"] = this.Lang
}
