# monitor-system
基于echo+xorm+mysql+redis+jwt+websocket+vue3+element-plus做的服务器云监控平台


# 简介
分布式云服务监控平台
1. 客户端收集服务器cpu、内存、磁盘等信息，通过websocket推送给服务端，服务端将信息推送给web进行面板可视化展示。当cpu、内存、磁盘超出预定阀值，右上角
提示报警信息。报警信息可手动忽略。
2. 账号列表是此平台的管理员集合。可以对此平台进行操作，如：添加服务检测、发布指令、升级客户端等
3. 服务检测服务于客户端所在的机器运行的程序。可以是Go、PHP、Java、Python等语言开发的业务程序。页面无需手动刷新，后台异步自动化调用，将服务检测的结果
可视化展示
4. 机器: 是客户端所在的服务器，可以是单台也可以是多台。客户端第一次部署，会生成一个唯一的客户端机器码。并提供页面列表输入框鼠标失去焦点标记备注信息。
服务端可以向某个客户端发送指令、升级客户端等。客户端升级保留最近三份历史记录（回滚暂未写）
5. 此平台管理员的所有操作将被记录到操作日志中收集
6. 服务端通过ssh连接客户端服务器


## 部署
创建服务端和客户端目录
```shell
mkdir /root/server && mkdir /root/client
```
### 本机编译
```shell
cd /server
./build.h

cd /client
./build.h
```

### 服务端

```shell
nohup ./server > nohup.out & echo $! > pidfile.txt
```

### 客户端
客户端注册成系统服务

```shell
./client install -d workDir

cd /etc/systemd/system && systemctl start cli-service.service
```

### web vue代码
[monitor-system-web](https://github.com/jeffcail/monitor-system-web)
