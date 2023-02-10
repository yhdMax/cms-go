FROM  golang:1.17-alpine
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/cms-vue-go

COPY  . $GOPATH/src/cms-vue-go
COPY ./config.json /etc/config.json

RUN go build -o cms-vue-go .

EXPOSE 8090

ENTRYPOINT ["./cms-vue-go"]
