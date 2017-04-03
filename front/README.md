前端测试文件夹

AJAX有跨域问题，`bee run`后前端人员在此写测试，模拟数据放在data

或者自己起一个apache或ngnix，指定一个域名或端口映射该文件夹，不需要`bee run`

如果`bee run`，打开浏览器`http://127.0.0.1:8080/front/index.html`开始测试