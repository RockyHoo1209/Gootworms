FROM golang:1.15

RUN mkdir /usr/Gootworms
WORKDIR /usr/Gootworms

COPY ./Gootworms ./Gootworms

RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.cn
# RUN go install -v ./...

# CMD cd /usr/work/