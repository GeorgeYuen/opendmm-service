FROM golang:1.9
MAINTAINER FandiYuan  <georgeyuan@diamondyuan.com>
copy / /opendmm/

RUN go get github.com/junzh0u/opendmm

RUN cd /opendmm && \
	go build