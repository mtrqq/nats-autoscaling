FROM golang:1.20-alpine as builder

WORKDIR /usr/app

COPY . /usr/app

RUN \
    --mount=type=cache,id=gomod,target=/go/pkg/mod \
    --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    go build -o consumer -ldflags "-s -w" ./consumer.go

FROM alpine:3.17

COPY --from=builder /usr/app/consumer /usr/bin

ENTRYPOINT ["consumer"]
