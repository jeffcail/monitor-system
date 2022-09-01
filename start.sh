#! /bin/bash

# nacos服务地址
nacos=192.168.0.125

if [ ! -f 'service-cloud-monitoring' ]; then
  echo 文件不存在! 待添加的安装包: 'service-cloud-monitoring'
  exit
fi

echo "service-cloud-monitoring..."
sleep 3
docker stop service-cloud-monitoring

sleep 2
docker rm service-cloud-monitoring

docker rmi service-cloud-monitoring
echo ""

echo "service-cloud-monitoring packing..."
sleep 3
docker build -t service-cloud-monitoring .
echo ""

echo "service-cloud-monitoring running..."
sleep 3

docker run \
  -p 9092:9092 \
  --name service-cloud-monitoring \
  -v /mnt/service-cloud-monitoring:/root/service-cloud-monitoring/log \
  -v /etc/localtime:/etc/localtime \
  -d service-cloud-monitoring \
  service-cloud-monitoring -ip $nacos -p 7848 -c service-cloud-monitor.yml


docker logs -f service-cloud-monitoring | sed '/Started service-cloud-monitoringApplication/q'

echo ""

