# README
   `TabakoAssistant` 项目是学习B站Up主 [墨明棋妙的豆子](https://space.bilibili.com/1532368659)相关教程后，实现的一个语音助手
  > 还在恶补前端知识，诸多功能未完善，仅作个人学习
## Features
### 后端
* 使用wails 框架，使用go v1.19 语言开发
### 前端
* 使用vue3 + vite2.0 + typescript 开发前端

## 功能
- [x] 托盘管理
- [X] 语音识别（需浏览器支持，自动开启）
- [X] 接入讯飞星火（v3.5）

## Todo
- [ ] 加入一些工具类，常见转码，如: base64转码
- [ ] 指定命令，如: 微信推送，B站查询，笔记录
  
## 常用命令
 * ### wails generate module  
  后端实现的函数，在前端调用，需要在[main.go](./main.go) 代码的bind里面传入相关对象，生成[wailsjs](./frontend/wailsjs)
 * ###   wails dev 
   热更新
 * ### wails build 
   构建app
   
  