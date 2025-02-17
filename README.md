# DeliciousTown



## 环境依赖
* golang 1.18
* mysql5.7
* redis
* nginx
* docker


## 项目简介
DeliciousTown是一个基于golang开发的一个开源的web网页文字游戏。
游戏主要分为两个部分，分别为游戏后台管理系统和Pc/H5网页游戏端，后台采用Golang+Vue+Mysql的前后端分离设计来实现游戏配置的管理。
游戏端主要采用Golang+Vue+Websocket+Mysql+Redis的前后端分离设计，Golang+Websocket主要承担游戏数据交互的任务，
例如：游戏聊天室、游戏公告推送、游戏数据实时更新等功能。
Golang+Vue前后端分离这种松耦合设计，可以有效的提高开发效率提高代码的可复用性可扩展性，通过前后端分离，可以更好地优化系统应用的性能和用户体验。
游戏数据交互采用Golang+Websocket这种设计可有效的提升程序并发与降低相关功能模块的开销，提升系统并发与流畅性，给玩家带来更好的游戏体验。
