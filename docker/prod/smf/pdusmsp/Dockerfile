FROM ubuntu:22.04

USER root

ENV http_proxy="http://203.91.199.136:80/"
ENV https_proxy="http://203.91.199.136:80/"
ENV no_proxy="localhost,127.0.0.1,192.168.0.0/16,172.16.27.15,172.16.27.22,172.16.27.23,10.209.184.15,10.209.184.19,10.209.184.23,172.20.0.0/16,10.96.0.0/12,172.16.27.0/24"

RUN apt-get update -y
RUN apt-get install iputils-ping -y
RUN apt-get install curl -y
RUN apt-get install telnet -y
RUN mkdir -p /opt/w5gc/bin
# Copy binary
COPY ./bin/pdusmsp /bin
#temporary
# RUN mkdir -p /home/requests
# COPY smContextCreate.json smContextRelease.json smContextUpdate.json n1msgtest n2infoext1test n2infotest /home/requests/

COPY ./bin/upfgw /bin


WORKDIR /opt/w5gc/bin
RUN mkdir -p /home/ubuntu/wipro5gc/testdata
# COPY testdata/n1msgtest testdata/n2infotest /home/ubuntu/wipro5gc/testdata
RUN export PATH=$PATH:/opt/w5gc/bin
# RUN export CONFIG_DIR=/root/go/src/w5gc.io/wipro5gcore/configs
#CMD rm -f /dev/shm/db /dev/shm/global_vm /dev/shm/vpe-api && \
#    exec pdusmsp && \
#    exec upfgw  
    # --request-type="update"

CMD ["/bin/sh", "-c", "rm -f /dev/shm/db /dev/shm/global_vm /dev/shm/vpe-api && exec pdusmsp & exec upfgw "]
