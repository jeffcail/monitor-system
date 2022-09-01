FROM alpine

# 设置工作目录
WORKDIR /root/service-cloud-monitoring

# 添加可执行文件
ADD ./service-cloud-monitoring $WORKDIR
#ADD dist /root/service-cloud-monitoring/dist

EXPOSE 9092

ENTRYPOINT ["./service-cloud-monitoring","-ip","192.168.0.125","-p","7848","-c","service-cloud-monitor.yml"]