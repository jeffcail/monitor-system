FROM alpine

# 设置工作目录
WORKDIR /root/client

# 添加可执行文件
ADD ./client $WORKDIR
#ADD dist /root/client/dist

EXPOSE 9092

ENTRYPOINT ["./client","-ip","192.168.0.125","-p","7848","-c","service-cloud-monitor.yml"]