# Multi Stages
# if golang base images not using alpine,disable c library using go library
# or add build tags(go build xxx -tags netgo or go build xxx -tags netcgo)
FROM golang:1.14.3-alpine3.11 as builder
# set up proxy
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build .


FROM alpine
WORKDIR /app
COPY --from=builder /app/rw-upload .
EXPOSE 9090
ENTRYPOINT ["./rw-upload"]