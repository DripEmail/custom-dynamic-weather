FROM golang:1.11-alpine3.8 as builder

WORKDIR /go/src/github.com/DripEmail/drip-personalized-weather

COPY . .

RUN apk update && \
    apk add --no-cache \
      wget \
      git && \
    rm -rf /var/cache/apk/* && \
    wget https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -O /bin/dep && \
    chmod +x /bin/dep && \
    dep ensure -vendor-only && \
    go build -o /drip-personalized-weather

FROM alpine:3.8

RUN apk update && \
    apk add --no-cache \
      ca-certificates \
      rm -rf /var/cache/apk/* \
    update-ca-certificates

COPY --from=builder drip-personalized-weather .

EXPOSE 8080

ENTRYPOINT ["./drip-personalized-weather"]
