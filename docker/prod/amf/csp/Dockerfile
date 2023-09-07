FROM ubuntu:22.04

USER root

ENV http_proxy="http://10.250.110.36:8000/"
ENV https_proxy="http://10.250.110.36:8000/"
ENV no_proxy="localhost,127.0.0.1,192.168.0.0/16"

RUN mkdir -p /opt/w5gc/bin
# Copy binary
COPY ./bin/csp /bin

#temporary
RUN mkdir -p /home/ubuntu/wipro5gc
COPY smContextCreate.json smContextRelease.json smContextUpdate.json n1msgtest n2infoext1test n2infotest /home/ubuntu/wipro5gc/

WORKDIR /opt/w5gc/bin
RUN export PATH=$PATH:/opt/w5gc/bin
# RUN export CONFIG_DIR=/root/go/src/w5gc.io/wipro5gcore/configs
CMD rm -f /dev/shm/db /dev/shm/global_vm /dev/shm/vpe-api && \
    exec csp --request-type="create"
