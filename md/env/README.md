### 搭建mysql + redis + nacos环境

```shell
chmod +x deploy_deocker.sh
chmod +x init_env.sh
```



### 安装

```shell
./deploy_deocker.sh
./init_env.sh



docker-compose -f yml/nacos-server.yml up -d
```

