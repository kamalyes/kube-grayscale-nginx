FROM golang:1.20 AS builder

WORKDIR /go/src
COPY ./src /go/src
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 go build -v -o main .

FROM alpine:3.19.0 AS api
RUN mkdir /app
COPY --from=builder /go/src/main /app
WORKDIR /app
ENTRYPOINT ["./main", "-v" ,"blue"]