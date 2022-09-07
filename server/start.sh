#! /bin/bash

# nacos服务地址
nacos=192.168.0.125

if [ ! -f 'server' ]; then
  echo 文件不存在! 待添加的安装包: 'server'
  exit
fi

echo "server..."
sleep 3
docker stop server

sleep 2
docker rm server

docker rmi server
echo ""

echo "server packing..."
sleep 3
docker build -t server .
echo ""

echo "server running..."
sleep 3

docker run \
  -p 9092:9092 \
  --name server \
  --net host \
  -v /mnt/server:/root/service-cloud-monitoring/server/log \
  -v /etc/localtime:/etc/localtime \
  -d server \
  server -ip $nacos -p 7848 -c service-cloud-monitor.yml

docker logs -f server | sed '/Started serverApplication/q'

echo ""