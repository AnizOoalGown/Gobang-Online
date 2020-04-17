# Gobang Online V0.1

##### 演示地址

[http://150.158.104.248:5555/](http://150.158.104.248:5555/)

##### 项目介绍

多人在线五子棋游戏测试版本v0.1，采用Vue, Gin, WebSocket, Redis等技术栈，前后端分离部署。用Chrome浏览器体验最佳，无账号登录，只支持无禁手规则。可打开多个网页登录多个玩家，一个玩家可同时进行或旁观多场游戏。

##### 功能介绍

1. 无账号用昵称登录
2. 大厅聊天
3. 房主创建房间
4. 挑战房主或旁观
5. 房间聊天
6. 开始游戏
7. 下棋、悔棋、投降、求和、逃跑
8. 游戏结束
9. 退出房间

##### 技术栈

前端采用vue框架，整合了：

1. element-ui：绘制页面布局、按钮、表格、输入框、弹框等组件
2. canvas：绘制棋子及棋盘
3. vuex：存储数据传输对象(DTO)，标签页信息等
4. vue-router：管理路由
5. vue-i18n：国际化
6. websocket：与后端通信

后端采用go语言的gin框架，整合了：

1. melody：websocket框架，管理session
2. redigo：设置连接池，访问redis
3. logrus：日志框架，整合mgorus，打印数据库到mongodb
4. viper：设置config.yml文件为配置文件
5. sync.RWMutex: 设置房间的锁，解决同步问题

中间件、数据库服务：

1. websocket：前后端通信，可参阅[API Doc.md](./API Doc.md)
2. redis：存储玩家信息、房间信息、大厅对话
3. mongodb：储存日志

部署：

​	腾讯云，CentOS 7，docker