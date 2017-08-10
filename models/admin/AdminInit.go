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
package admin

// still keep chinese some...

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hunterhug/GoWeb/lib"
	"github.com/hunterhug/GoWeb/models/blog"
	"math/rand"
)

func InitData() {
	InsertUser()
	InsertGroup()
	InsertRole()
	InsertNodes()
	InsertConfig()
	InsertCategory()
	InsertRoll()
	InsertPaper()
}

//插入网站配置
func InsertConfig() {
	fmt.Println("insert config start")
	c := new(blog.Config)
	c.Photo = "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
	c.Id = 1
	c.Title = "Rabbit(Tuzi) Enterprise Web"
	c.Webinfo = `
	{
	"1":{"name":"About","limit":6},
        "2":{"name":"News","limit":6},
        "3":{"name":"Lifes","limit":6},
        "4":{"name":"Production","limit":6},
        "5":{"name":"Flower","limit":6},
        "6":{"name":"TeaCup","limit":6}
	}
	`
	c.Phone = "0750-12345678"
	c.Content = `
<div align="center">
	<p>
		<img src="/file/image/53/68756e7465726875671e5573ac53bb5813b6b51d47d2db806b.gif" alt="" />
	</p>
</div>
	`
	c.Slogan = "A Enterprise Web, You can have a try"
	c.Address = `<meta description="GoWeb" >
<!-- some other script put in here -->`
	c.Code3 = `Power by hunterhug at 2017`
	c.Code2 = `统计代码`
	c.Code1 = "跟帖代码"
	err := c.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert config end")
}

func InsertCategory() {
	fmt.Println("insert category start")
	cs := map[int64]string{1: "About", 2: "News", 3: "Lifes", 4: "Production", 5: "Flower", 6: "TeaCup"}
	for k, v := range cs {
		c := new(blog.Category)
		c.Id = k
		c.Title = v
		c.Createtime = lib.GetTime()
		c.Status = 1
		c.Image = "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
		c.Content = v
		if k > 4 {
			c.Pid = 4
		}
		err := c.Insert()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("insert category end")
}

func InsertPaper() {
	k := 40
	for k > 0 {
		k = k - 1
		paper := new(blog.Paper)
		paper.Title = "Test test test test data"
		paper.Status = 1
		paper.Photo = "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
		paper.Descontent = ""
		paper.Content = `
毌唋芩郫拶灟椼裶机汊庤忭坳炄岏祹屮芮髧耡。郰一浀牼棫痷丌葂謔夬，倕兀秎紩湨瘏乜蓖夃仜。惎泐朸瞝嗂諲扡一硰訞悢竻尒宬屮，葸傒乇冞砓誆怞匜亍呫瘑垕仂枑。跍屇匟凞滁鍺玎一棌欳桏岵汃玵丌，赨詎丌忥訇搨枍帄亍岥頖砒夃泔。偝一斪郬堬傿丌漰聬丏，猏亍峉莩堜筤屮銆旡氕。湸枘厊篹蒆麮艼一堨蚺晑攽阞眄屮，葧湥丌怭挌犎祊戉屮岮銠笀爿狌。媊盰仵澮幎鴗仴一喦梪倢咁庀拺乜，慀揝乜侕茇頏呺犰丌玝漵畇仉呯。湫咑扞諤嵥犝匟一焨娸淯怲忉苶兀，楱嵙丌牬苻嵞抴庂屮迋颮拶气竻。
淭一厘晢臹隒丌蒛霒冘，庳乜砐淈琲葙丌漍厹刌。仈茖呴釸柟襻頎獊夼匢挀杙枎侞牣椲亍泐熁搒。仈迼枘笘恉鷿跫睔冱犴玴庉徂泑耴葨乇虰餀粴。裀一咺偝喓葖屮膃璔卬，淉兀姷觖嵧絻乇馺巿犮。娹一炰焌椋馵丌鄚澣仈，舺亍柊捵殔裺丌摥旡戉。淝一咼晢愩兀蜛駬丱，莐亍瓮悺愊媸兀翞仈丼。釱一俓祪鄆溱屮幓樴旡，谻屮峆焄喒媺丌僬仂忉。晜一衎悾趄搎乜綮醍亓，豝丌峟梩窙寖屮銗丱仡。
豜一盄欸毰裾亍蜠鋿丱，莣亍疪婟愝鉖乜蒚仂仝。仉郟佸秺籺齶跠麀匟朸芔辿阽刲刞痭乜狌豨蛖。眭一眅郫覘滒兀塿踳旡，掔兀袀晡棜傸兀潳仈仜。揌阹芀橯瑑擖异一斮紩挩炖刉峛乇，腞軮亍泞奅酯佌氻兀祌蝂垛殳芠。确姖屻磣裲瞝夼一堩崌粄邴氶咭屮，蛖軦乜呡胅鉏厔仜屮忞馺卼冘芡。逋一狪涾渃趌亍蒠橛殳，蚿屮柅袤厧筭屮賏冘帄。寎厒汆燛萰撉妅一莿淴剞玢忉珋乜，覛睎乜咍紃巰芫氿亍芞摬俍丮呥。
菡忞旮踽歅糒忔一軩崦唅枟氻朐兀，酯梴丌劼拹鉏沶刉丌炅榹笀亓坶。殳珜狋勖咾鸆蒍綍忕扞姞岍怚怳忮睧亍杻鳲葨。偨邯彴蝹骭獪奷一欹荵祏姖刉胣亍，筴焟丌妵狦搹岟夗兀糽輐拏仂呧。气秜狖琇苹鱢蒆隗圴刐胉阯咘怐沄鈲乇佶誏楒。圠茛沰桷洠韥艉絺伀朻袀邞迕戔妎葌屮弤魡睧。媢怓伈膦艉瘲扙一掱崰笄杸仩虴屮，羦湀亍抳眊傱矸屴兀坫煻苶夃彔。翏一姝淗朁蜍屮瞅撉尐，婑丌浂秺粞溞丌僝夬仨。粣一虷蚹椆觜乇瞁螅丮，捸丌挀梖軺溲兀榯仈扐。
塈昅刓窵裶黖伒一腏桼豗呣尒秎屮，葹嵉乜岦扂滈玢氕亍枍戫柁厹怳。卬埇牪釬砒齶綈愲汔囟迮忮沝刲厎塥乇枑銣痶。逯毞幵獪趏澨仱一敧埮毤芠匜袃丌，媿掰丌坵砅稘抰仝亍呥馜甮殳厒。夬栲矻痎咼鱢滀朠扚厊峏灴佽沝犺雎屮怍滹搳。烶一侳淟酢盝兀榬蕠尐，畣屮勀娹旒嗈乜殠圠犰。亓唋炖梐苲齶鄏貅扚伄勂抏极弤佟獊屮昄熐覛。夬軑泒捼奅灝暕煒汆庄炡肐徂拑阭塯丌怌僓煰。
	`
		paper.Author = "hunterhug"
		paper.Createtime = lib.GetTime()
		paper.Cid = int64(rand.Intn(6) + 1)
		paper.Istop = int64(rand.Intn(2))
		paper.Insert()
	}
}

func InsertRoll() {
	rolls := map[string]string{
		"tuzi":   "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"me":     "/file/image/37/68756e7465726875673308fd68c821f8fb4180732625ef10ba.png",
		"tuzizi": "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
	}
	for k, v := range rolls {
		t := new(blog.Roll)
		t.Photo = v
		t.Status = 1
		t.Title = k
		t.Createtime = lib.GetTime()
		t.Insert()
	}
}

// 用户数据
func InsertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.Username = beego.AppConfig.String("rbac_admin_user")
	u.Nickname = "TuziAdmin"
	u.Password = lib.Pwdhash(beego.AppConfig.String("rbac_admin_user"))
	u.Email = "569929309@qq.com"
	u.Remark = "God in Rabbit Country"
	// 2 stand for close, but it has very high authority
	u.Status = 2
	u.Createtime = lib.GetTime()
	err := u.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	u1 := new(User)
	u1.Username = "test"
	u1.Nickname = "TuziTest"
	u1.Password = lib.Pwdhash("test")
	u1.Email = "569929309@qq.com"
	u1.Remark = "Just a Test User"
	u1.Status = 2
	u1.Createtime = lib.GetTime()
	err1 := u1.Insert()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert user end")
}

// 模组数据
func InsertGroup() {
	fmt.Println("insert group ...")
	g := new(Group)
	g.Name = "兔子后台"
	g.Title = "后台管理"
	g.Sort = 1
	g.Id = 1
	g.Status = 1
	e := g.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	g1 := new(Group)
	g1.Name = "兔子后台"
	g1.Title = "文章管理"
	g1.Sort = 2
	g1.Id = 2
	g1.Status = 1
	e = g1.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	g2 := new(Group)
	g2.Name = "兔子后台"
	g2.Title = "图片管理"
	g2.Sort = 3
	g2.Id = 3
	g2.Status = 1
	e = g2.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println("insert group end")
}

// 角色数据
func InsertRole() {
	fmt.Println("insert role ...")
	r := new(Role)
	r.Name = "管理员"
	r.Remark = "权限最高的一群人"
	r.Status = 1
	r.Title = "管理员角色"
	r.Insert()
	fmt.Println("insert role end")
}

// 节点数据
func InsertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	g1 := new(Group)
	g1.Id = 2
	g2 := new(Group)
	g2.Id = 3
	nodes := []Node{
		/*

			RBAC管理中心

		*/
		{Id: 1, Name: "rbac", Title: "权限中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		{Id: 2, Name: "node/index", Title: "节点管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 3, Name: "Index", Title: "节点首页", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 4, Name: "AddAndEdit", Title: "增编节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 5, Name: "DelNode", Title: "删除节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},

		{Id: 6, Name: "user/index", Title: "用户管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 7, Name: "Index", Title: "用户首页", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 8, Name: "AddUser", Title: "增加用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 9, Name: "UpdateUser", Title: "更新用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 10, Name: "DelUser", Title: "删除用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},

		{Id: 11, Name: "group/index", Title: "分组管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 12, Name: "Index", Title: "分组首页", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 13, Name: "AddGroup", Title: "增加分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 14, Name: "UpdateGroup", Title: "更新分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 15, Name: "DelGroup", Title: "删除分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},

		{Id: 16, Name: "role/index", Title: "角色管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 17, Name: "index", Title: "角色首页", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 18, Name: "AddAndEdit", Title: "增编角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 19, Name: "DelRole", Title: "删除角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 20, Name: "GetList", Title: "列出角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 21, Name: "AccessToNode", Title: "显示权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 22, Name: "AddAccess", Title: "增加权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 23, Name: "RoleToUserList", Title: "列出角色下用户", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 24, Name: "AddRoleToUser", Title: "授予用户角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},

		/*

			配置中心

		*/
		{Id: 25, Name: "config", Title: "配置中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		//-------
		//网站配置
		{Id: 26, Name: "option/index", Title: "网站配置", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 27, Name: "Index", Title: "网站配置首页", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		{Id: 28, Name: "UpdateOption", Title: "更新网站配置", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		//网站配置
		//个人信息
		{Id: 29, Name: "user/index", Title: "个人信息", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 30, Name: "Index", Title: "个人信息首页", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		{Id: 31, Name: "UpdateUser", Title: "更新个人信息", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		//个人信息

		/*

			文章中心

		*/
		{Id: 32, Name: "blog", Title: "文章中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g1},
		//------
		//文章目录
		{Id: 33, Name: "category/index", Title: "目录列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 34, Name: "Index", Title: "目录列表首页", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 35, Name: "AddCategory", Title: "增加目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 36, Name: "UpdateCategory", Title: "修改目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		//文章目录
		//文章
		{Id: 37, Name: "paper/index", Title: "文章列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 38, Name: "Index", Title: "文章列表首页", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 39, Name: "AddPaper", Title: "增加文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 40, Name: "UpdatePaper", Title: "修改文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 41, Name: "DeletePaper", Title: "回收文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 42, Name: "RealDelPaper", Title: "删除文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		//文章

		/*

			图片管理

		*/
		{Id: 43, Name: "picture", Title: "图片中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g2},
		//---------
		//相册
		{Id: 44, Name: "album/index", Title: "相册管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 45, Name: "Index", Title: "相册首页", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 46, Name: "AddAlbum", Title: "增加相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 47, Name: "DeleteAlbum", Title: "删除相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 48, Name: "UpdateAlbum", Title: "修改相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		//相册
		//图片
		{Id: 49, Name: "photo/index", Title: "图片管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 50, Name: "Index", Title: "图片首页", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 51, Name: "AddPhoto", Title: "增加图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 52, Name: "DeletePhoto", Title: "回收图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 53, Name: "UpdatePhoto", Title: "修改图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 54, Name: "RealDelPhoto", Title: "删除图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		//图片

		//补充的
		{Id: 55, Name: "DeleteCategory", Title: "删除目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},

		{Id: 56, Name: "paper/rubbish", Title: "文章回收站", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 57, Name: "Rubbish", Title: "文章回收站", Remark: "", Level: 3, Pid: 56, Status: 1, Group: g1},

		{Id: 58, Name: "photo/rubbish", Title: "图片回收站", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 59, Name: "Rubbish", Title: "图片回收站", Remark: "", Level: 3, Pid: 58, Status: 1, Group: g2},

		//首页图片轮转
		{Id: 60, Name: "roll/index", Title: "首页轮转", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 61, Name: "Index", Title: "轮转列表", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 62, Name: "AddRoll", Title: "增加轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 63, Name: "UpdateRoll", Title: "更新轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 64, Name: "DeleteRoll", Title: "删除轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
	}
	for _, v := range nodes {
		n := new(Node)
		n.Id = v.Id //这句是无效的,后来 bug 被 beego 官方改好了
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		e := n.Insert()
		if e != nil {
			fmt.Printf("%#v:%#v\n", n, e.Error())
		}
	}
	fmt.Println("insert node end")
}
