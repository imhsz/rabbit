# Golang网站手脚架

项目代号： 兔子(tuzi)

![兔子](tuzi.png)

实现了基本的RBAC模块和博客模块（文章和相册），附带Dashbord后端UI，基本框架形成，依靠此项目可敏捷开发。

1. 基于角色的访问控制（Role-Based Access Control）作为传统访问控制
2. Amaze UI v2.7.0（部分后台）和jQuery EasyUI 1.4.2（后台表格CRUD）、Bootstrap（前台）混合
3. 前台页面大量使用Vue.js，因为只查询的前台路由全部返回JSON。（待做)

## 使用说明
使用只需拉下库

```
go get -v github.com/hunterhug/GoWeb
```

编译程序

```
go build
```

初始化数据库

```
./GoWeb -initdb
```

运行程序

```
./GoWeb
```

启动前请配置conf/app.conf中的数据库

```
appname = tuzi
version = 1.0.0

# 生产环境改为prod
runmode = dev

###################

# 可以直接通过静态访问的文件夹，位于根目录下面
StaticDir = static:static file:file

# 国际化语言
lang_types = en-US|zh-CN

# 路由区分大小写
RouterCaseSensitive = false

# 中国时间请设为8，不然数据库时间会混乱
timezone = 8

# 调试数据库 close/open
dblog = close

###################

# 前台模板，可以改,wordpress功能
home_template = home/default
admin_template = admin/default

# 文件上传保存地址，后面不可以是/，必须是根目录下的文件夹，为了速度更快，文件直接到前端，可改写
filebasepath = file

###################

# 权限控制，建议不要乱改
sessionon = true
sessionname = beautysessionid
sessionhashkey = mostbeautyart
rbac_role_table = role
rbac_node_table = node
rbac_group_table = group
rbac_user_table = user
rbac_admin_user = admin
not_auth_package = public,static,home,file

###################

# 0不验证，1验证，2实时验证,建议不要改
user_auth_type = 2
rbac_auth_gateway = /public/login

# cookie一周内登录开关，1表示开，建议设为0
cookie7 = 0

[dev]
httpport = 8080
db_host = 127.0.0.1
db_port = 3306
db_user = root
db_pass = root
db_name = tuzi
db_type = mysql
db_prefix = tb_

[prod]
EnableGzip = true
httpport = 80
db_host = 127.0.0.1
db_port = 3306
db_user = root
db_pass = root
db_name = tuzi
db_type = mysql
db_prefix = tb_
```

后台入口：http://127.0.0.1:8080/public/login

账户admin，密码：admin

![](doc/login.png)
![](doc/admin.png)

## 开发流程
文件结构
```
    ----conf 配置文件夹

        ----app.conf 		应用配置文件
        ----local_**.ini 	国际化文件

    ----controllers 控制器
        ----admin	后台控制器
            ----blog 博客模块
            ----rbac 权限模块
        ----home 	前台控制器
        ----rbac.go 路由权限过滤器入口

    -----lib 公共库
    -----file 上传文件保存地址
    -----models ORM模型
        ----admin RBAC主要数据库
            ----AdminInit.go 数据默认填充地
        ----blog  博客主要数据库

    ----routers 路由
    ----static  静态文件
    ----views	视图
        ----admin 	后台视图
            ----default 默认主题
        ----home 	前台视图
            ----default 默认主题

    ----log 日志
    ----doc 说明文档（重点，待写）
    ----test 测试文件夹
```

>RBAC权限相关的models统一放在admin文件夹，其他都放在home文件夹.
	前台控制相关的controllers统一放在home文件夹，其他都放在admin文件夹
	URL router统一M/C/A方式，该正则url需要验证权限，如rbac/public/index（三个参数），其他如public/index不验证。

>登录说明
	登陆过的用户只能注销后登录，支持定义cookie登录。进入后台时验证session，session不存在则验证cookie，如果用户未被冻结，增加session，
	同时更改用户登录时间、登录IP等，cookie与登录IP绑定（防止cookie劫持）。

>系统时间默认数据库本地时间为东八区，北京时间。

>后台模板在views/admin，前台模板在views/home，子文件夹为主题，默认主题为default

>所有配置在conf文件夹conf/app.conf，支持国际化

>数据库数据填充在models/*/*Init.go中定义

>视图模板均放在static中

>前台首页配置（可动态调整首页）

```
{"1":{"name":"每日动态","limit":6},
"2":{"name":"画室动态","limit":6},
"3":{"name":"招生动态","limit":6},
"4":{"name":"美术资讯","limit":6},
"5":{"name":"高考喜报","limit":6},
"6":{"name":"学员风采","limit":3},
"7":{"name":"教师风采","limit":3},
"8":{"name":"学生作品","limit":6}}
```

每次在models/admin/AdminInit.go增加路由权限，请执行

```
./GoWeb -rbac
```

## Ngnix架站

配置 nginx.conf,，server_name为域名，access_log为日志路径（要手动建文件夹）

```
server{
        listen 80;
        server_name beauty.lenggirl.com www.beautyart.top;
        charset utf-8;
        access_log /data/logs/nginx/beauty.lenggirl.com.log;
        #error_log /data/logs/nginx/www.lenggirl.com.err;
        location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_pass http://localhost:8080;
	    proxy_set_header X-Real-Ip $remote_addr;
        }

}
```

## 计划（2017.4-5）
1. 前台重建
2. 逻辑优化

