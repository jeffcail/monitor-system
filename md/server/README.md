### 1.搭建mysql、Redis

### 2. 新建 monitor数据库导入monitor.sql文件

### 3. 配置文件

```ini
修改database 和 redis 节点信息
DbDsn
RedisAddr && Password
```

### 4. 编译服务端可执行文件

```shell
./build.sh
```

### 5. 部署服务端

```shell
1. cd /root && mkdir server

2. 将dist目录上传至 server目录下

3. 将conf目录上传至 server目录下

3. nohup ./server > nohup.out & echo $! > pidfile.txt
```

### 6. 访问服务端控制台 

**url:9999**

账号: admin
密码: 123456