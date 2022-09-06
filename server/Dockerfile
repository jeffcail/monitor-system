FROM alpine

# 设置工作目录
WORKDIR /root/service-cloud-monitoring/server
# 添加可执行文件
ADD ./server $WORKDIR
ADD dist /root/service-cloud-monitoring/server/dist

EXPOSE 9092

ENTRYPOINT ["./server","-ip","192.168.0.125","-p","7848","-c","service-cloud-monitor.yml"]
