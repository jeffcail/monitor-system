### 1. 编译客户端

```shell
./build.sh
```


### 2. 部署客户端

```SHELL
1. cd /root && mkdir client

2. 修改 conf/conf.ini 配置文件 GoFileServe:  ip:port（ip: 服务端的ip:服务端端口）

3. 将客户端的conf目录上传至client 目录下

4. 将编译后的客户端可执行文件上传至 /root/client目录下
```


### 3. 安装客户端服务

```shell
cd /root/client

./client install
```


### 4. 启动客户端服务

```shell
cd /etc/systemd/system && systemctl start client-monitor.service
```

