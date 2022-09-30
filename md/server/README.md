### 1.搭建mysql、Redis 和 nacos

### 2. 新建 monitor数据库导入monitor.sql文件

### 3. nacos新增配置

```markdown
Data ID: service-cloud-monitor.yml
Group: monitor


修改 配置文件 MYSQL:        DbDsn
修改 配置文件 Redis:        RedisAddr && Password
修改 配置文件 GoFileServe:  ip:port（ip: 服务端的ip:服务端端口）
```

### 4. 编译服务端可执行文件

```shell
./build.sh
```

### 5. 部署服务端

```shell
1. cd /root && mkdir server

2. 将dist目录上传至 server目录下

3. nohup ./server > nohup.out & echo $! > pidfile.txt
```

### 6. 访问服务端控制台 

**url:9092**

账号: admin
密码: 123456