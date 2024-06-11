# hmdp-go
* [参考仓库地址](https://hmdp-go)

最近在b站学习redis[黑马程序员Redis入门到实战教程，深度透析redis底层原理+redis分布式锁+企业解决方案+黑马点评实战项目](https://www.bilibili.com/video/BV1cr4y1671t?p=25&vd_source=d4c7ab0b2bf2d254f64ccd5ee3bbbdfd)
学习到实战篇时发现语言是Java代码,但我又想用go语言学习,遂参照hmdp项目写下了这个go语言实现的hmdp项目demo版

* 该项目是gin+vue的前后端分离项目，使用gorm访问MySQL

* 项目结构进行分层，使用依赖注入的方式对项目进行解耦---[Gin实现依赖注入教程](https://www.cnblogs.com/FireworksEasyCool/p/11805148.html)

* 使用jwt，对API接口进行权限控制---[gin-jwt对API进行权限控制教程](https://www.cnblogs.com/FireworksEasyCool/p/11455834.html)

* 使用[go-playground/validator](https://github.com/go-playground/validator)开源库简化gin的请求校验---[gin请求数据校验教程](https://www.cnblogs.com/FireworksEasyCool/p/12794311.html)

* 在token过期后的半个小时内，用户再次操作会自动刷新token

### 项目结构

<pre><code>
├── cmd  程序入口
├── common 通用模块代码
├── config 配置文件
├── controller API控制器
├── sql 数据库文件
├── models 数据表实体
├── page 页面数据返回实体
├── repository 数据访问层
├── router 路由
├── service 业务逻辑层
├── vue-admin Vue前端页面代码
</code></pre>

### 下载安装项目
`go get -x hmdp-go/cmd`

### go后台程序运行方式

1.在MySQL中运行文件夹/docs中的mysql.sql脚本

2.在gin-vue-admin/cmd目录下运行`go run main.go`

### vue前端运行方式

请看文件夹/vue-admin中的README.md
