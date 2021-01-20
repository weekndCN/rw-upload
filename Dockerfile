# Multi Stages
FROM golang:1.14 as builder
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