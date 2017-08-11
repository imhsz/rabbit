package home

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/GoWeb/models/blog"
	//"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/hunterhug/GoWeb/lib"
)

type MainController struct {
	baseController
}

//首页配置定义放哪些东西
type paperablum map[string]map[string]interface{}

//配置
var config *blog.Config

func (this *MainController) Prepare() {
	this.baseController.Prepare()
	config = new(blog.Config)
	config.Id = 1
	config.Read()
	//网站配置
	this.Data["config"] = config
	this.Data["category"] = getmulu(0, 0)
	this.Data["photo"] = getmulu(0, 1)

}

func (this *MainController) Index() {
	a := this.GetString("lang", "")
	if a != "" {
		this.Ctx.SetCookie("lang", a)
		switch a {
		case "en":
			this.Lang = "en-US"
		case "cn":
			this.Lang = "zh-CN"
		default:
			this.Lang = "zh-CN"
		}
		this.Data["Lang"] = this.Lang
	}
	//轮转图
	roll := new(blog.Roll)
	rolls := []orm.Params{}
	roll.Query().Filter("Status", 1).OrderBy("-Sort", "Createtime").Values(&rolls)
	this.Data["roll"] = rolls

	//首页
	index := paperablum{}
	//println(config.Webinfo)
	err := json.Unmarshal([]byte(lib.TripAll(config.Webinfo)), &index)
	if err != nil {
		beego.Trace(err.Error())
	}
	//beego.Trace("%v",index)

	for i, item := range index {
		_, td, tc := getjinhan(item["name"].(string), int(item["limit"].(float64)))
		//beego.Trace("%v", tc)
		this.Data["t"+i] = td
		this.Data["t"+i+"c"] = tc
	}

	//beego.Trace(this.Data["Lang"])
	this.TplName = this.GetTemplate() + "/index.html"
}

func getmulu(beautyid int, blogtype int) []orm.Params {
	//目录
	//文章列表首页
	category := new(blog.Category)
	categorys := []orm.Params{}
	//查询条件：缀美文章类型，一级
	category.Query().Filter("Status", 1).Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys, "Id", "Title")
	for _, cate := range categorys {
		//二级
		son := []orm.Params{}
		category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "Createtime").Values(&son, "Id", "Title")
		cate["Son"] = son
	}
	return categorys

}

func getjinhan(title string, count int) (error, []orm.Params, orm.Params) {
	err, category := getcategory(title)
	if err != nil {
		err, album := getalbum(title)
		if err != nil {
			return errors.New("找不到该目录"), []orm.Params{}, album
		} else {
			id := album["Id"].(int64)
			return nil, getphoto(id, count), album
		}
	} else {
		id := category["Id"].(int64)
		return nil, getpaper(id, count), category
	}
}

//获取开启的文章，按置顶
func getpaper(id int64, count int) []orm.Params {
	paper := new(blog.Paper)
	papers := []orm.Params{}
	paper.Query().Filter("Cid", id).Filter("Type", 0).Filter("Status", 1).OrderBy("-Istop", "Createtime").Limit(count, 0).Values(&papers)
	return papers
}

//获取开启的图片，按轮转，置顶
func getphoto(id int64, count int) []orm.Params {
	paper := new(blog.Paper)
	papers := []orm.Params{}
	paper.Query().Filter("Cid", id).Filter("Type", 1).Filter("Status", 1).OrderBy("-Isroll", "-Istop", "Createtime").Limit(count, 0).Values(&papers)
	return papers
}

//获取文章目录
func getcategory(title string) (error, orm.Params) {
	category := new(blog.Category)
	categorys := []orm.Params{}
	category.Query().Filter("Type", 0).Filter("Siteid", 0).Filter("Title", title).Limit(1).Values(&categorys)
	if len(categorys) == 0 {
		return errors.New("找不到文章分类"), orm.Params{}
	} else {
		return nil, categorys[0]
	}
}

//获取相册目录
func getalbum(title string) (error, orm.Params) {
	category := new(blog.Category)
	categorys := []orm.Params{}
	category.Query().Filter("Type", 1).Filter("Siteid", 0).Filter("Title", title).Limit(1).Values(&categorys)
	if len(categorys) == 0 {
		return errors.New("找不到相册分类"), orm.Params{}
	} else {
		return nil, categorys[0]
	}
}
