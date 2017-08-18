# Project： Rabbit(Tuzi)

兔子脚手架是用Golang语言开发的项目，主要用来快速开发企业网站。

感谢`beego`,`jquery`,`easyui`,`bootstrap`等框架的贡献者～

[中文介绍](doc/Chinese.md) anglicizing...

You can own a enterprise web just listen to me! You just need install golang environment(ask for google help).Under developing...

![Rabbit](tuzi.png)

## 1. How to use

Just do this:

```shell
go get -v github.com/hunterhug/GoWeb
```

Or

```
git clone https://www.github.com/hunterhug/GoWeb
mkdir %GOPATH%/src/github.com/hunterhug
mv GoWeb %GOPATH%/src/github.com/hunterhug
```

Then build our web

```shell
go build
```

Before run, Please config the db set in `conf/app.conf`, use Mysql(install can ask google help)

```
# you can set it into prod when in production environment
runmode = dev

[dev]
httpport = 8080
db_host = 127.0.0.1
db_port = 3306
db_user = root
db_pass = root
db_name = tuzi
db_type = mysql

[prod]
EnableGzip = true
httpport = 80
db_host = 127.0.0.1
db_port = 3306
db_user = root
db_pass = root
db_name = tuzi
db_type = mysql
```

And init our database

```shell
./GoWeb -db=1
```

~~Or(I recommend to use)~~

```
cd doc
cd sh
./initdb.sh tuzi
```

`tuzi` is your db name, script equal to `mysql -uroot -p -v tuzi < init.sql`

Last run it

```shell
./GoWeb
```

Ok, you can open `http://127.0.0.1:8080`

Login to edit the website: `http://127.0.0.1:8080/public/login`

User: `admin`

Password：`admin`

if upload file error please make a new dir names `file` under this project: 

```
# if in linux
mkdir file
chmod 777 file
```

## 2. How to Develop

### a. Project Structure(modularization)

```shell
    ----conf config module

        ----app.conf 		app config file
        ----local_**.ini 	internationalization file

    ----controllers   controllers module
        ----admin	  back-end
            ----blog  blog-edit(category/paper) module
            ----rbac  authority module
        ----home 	front end
        ----rbac.go router authority filtering

    -----lib  public Library
    -----file upload file keep in here
    
    -----models ORM module
        ----admin RBAC database operation
            ----AdminInit.go admin data fill by this
        ----blog  blog database operation

    ----routers url router
    ----static  static file such as css/js
        ---admin  back-end js/css
        ---home  front-end js/css
            ---amazi  Meizi UI(China)
            ---boostrap Most Niubi UI
         ---tool some tool js
         ---diy our diy js/css
    ----views	 template views
        ----admin 	back-end
            ----default defaule theme
        ----home 	front-end
            ----default default theme

    ----front can use for vue/angular... preparing
        ---data JSON data

    ---help  help yu init db
        --- init.sql important data
        --- ngnix-tuzi.conf Nginx config
```

We have already implement basic RBAC module and Blog module（Article and Album equal to enterprise News and Production）, And have a Dashboard back-end UI, The UI can accelerated development. 
which inspiration by：[http://www.beautyart.top](http://www.beautyart.top), you can visit it to see how it is.

1. Role-Based Access Control
2. Amaze UI v2.7.0（little back-end）和jQuery EasyUI v1.4.2（back-end table CRUD）、Bootstrap v3.3.5（front-end）mixed
3. Prepare use Vue.js v2.2.6 to separate back-end and front-end（Maybe）, back-end just offer REST JSON API, and front-end can first test Off-line then if no problem, docking!
when ajax call JSON must pay attention across-domain rule(see rht dir front), why use this way due to can reduce the back-end burden~~ and more fast develop...

### b. Rules And Explanation

1. RBAC function must put in `controllers/admin` folder.Front-end controllers put in `controllers/home` folder, other put in `controllers/admin`.URL router use `M/C/A` ways, such router  `rbac/public/index`（three）must authorize.
2. Login：you can logout after login, support cookie remember login. when enter back-end, check session, if not exist session then check cookie. if user is activated, add the session, record login times、login IP etcd. When remember login, will add cookie(cooke bind by ip and encrypted password for hijacking prevention).
3. System time default timezone UTC/GMT+08:00 China BeiJin, you can change in `app.conf`.
4. Back-end template in `views/admin`, front-end in `views/home`, the sub folder is theme, which default is default... can change in `app.conf`.
5. All config in `conf/app.conf`, support internationalization, can use chinese/english by browser such `Accept-Language:en-US,en;q=0.5` 
6. All data initialization can define in `models/*/*Init.go`, I will change it all in english.
7. All js/css such static file must put in `static`
8. Website home can be change by this（just ignore it...waiting for explanation）

```
{
"1":{"name":"about","limit":6}
}
```

### c. Add routers and permissions

Every add routers and permissions in `models/admin/AdminInit.go`, please rebuild rbac:

```
./GoWeb -rbac=1
```

debug you can use `bee run`...

## 3. How to use Nginx(optional)

First install Nginx(ask for google...)

Enter `/usr/local/nginx/conf`

```
vim nginx.conf
```

In the last of `nginx.conf`, add:

```
include sites/*.conf;
```

New a `sites` folder， put our `help/ngnix-tuzi.conf` under it：

config `ngnix-tuzi.conf`, `server_name` is your domain, `access_log` is the log path（you must makedir first）

```shell
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

# Have a Look!

![](doc/img/index.png)

![](doc/img/blog.png)

# LICENSE

```
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
```

Welcome Add PR/issues.

For questions, please email: gdccmcm14@live.com.

