# 📺 抖音弹幕监听器

## 😎介绍及配置

### 介绍

基于系统代理抓包打造的抖音弹幕服务推送程序，它能够获取浏览器直播间上抖音弹幕数据，它可以监听**弹幕**，**点赞**，**送礼**，**进入直播间**，等系列消息，你可使用它做自己的直播间数据分析，以及弹幕互动游戏，语音播报等。


### 推送数据格式

弹幕数据由WebSocket服务进行分发，使用Json格式进行推送，可前往[ws在线测试](http://wstool.jackxiang.com/)网站，连接 ws://127.0.0.1:8709/ws 进行测试

### 监听控制
通过HTTP服务API进行控制，POST请求 http://127.0.0.1:8709/api
参数form-data
- roomId:42643374680 (web房间号)
- command:close （connect表示监听、close表示取消监听）

### 使用方法
docker run -p 8709:8709 javpower/dy:v1.0.0

## ⚖️免责声明

+ 本程序仅供学习参考，不得用于商业用途，不得用于恶意搜集他人直播间用户信息!
