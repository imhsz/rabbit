# Project： Rabbit(Tuzi)

[中文介绍](Chinese.md)

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

Or(I recommend to use)

```
cd help
./initdb.sh tuzi
```

`tuzi` is your db name

Last run it

```shell
./GoWeb
```

Ok, you can open `http://127.0.0.1:8080`

Enter to edit the website: `http://127.0.0.1:8080/public/login`

User: `admin`

Password：`admin`

if upload file error please make a new dir names `file` under this project: 

```
# if in linux
mkdir file
chmod 777 file
```

# LICENSE

```
Copyright 2017 by GoSpider author.
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

