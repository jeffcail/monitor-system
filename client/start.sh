#! /bin/bash

# nacos服务地址
nacos=192.168.0.125

if [ ! -f 'client' ]; then
  echo 文件不存在! 待添加的安装包: 'client'
  exit
fi

echo "client..."
sleep 3
docker stop client

sleep 2
docker rm client

docker rmi client
echo ""

echo "client packing..."
sleep 3
docker build -t client .
echo ""

echo "client running..."
sleep 3

docker run \
  -p 9092:9092 \
  --name client \
  -v /mnt/client:/root/client/log \
  -v /etc/localtime:/etc/localtime \
  -d client \
  client -ip $nacos -p 7848 -c service-cloud-monitor.yml


docker logs -f client | sed '/Started clientApplication/q'

echo ""

