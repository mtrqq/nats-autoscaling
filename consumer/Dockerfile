FROM golang:1.20-alpine as builder

WORKDIR /usr/app

COPY . /usr/app

RUN \
    go mod download && \
    go build -o consumer -ldflags "-s -w" ./consumer.go

FROM alpine:3.17

COPY --from=builder /usr/app/consumer /usr/bin

ENTRYPOINT ["consumer"]
