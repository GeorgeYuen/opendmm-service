FROM golang:1.9-alpine
MAINTAINER FandiYuan  <georgeyuan@diamondyuan.com>
copy / /opendmm/

RUN go get github.com/junzh0u/opendmm

RUN cd /opendmm && \
	go build