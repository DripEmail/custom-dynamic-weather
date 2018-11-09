FROM golang:1.11-alpine3.8 as builder

WORKDIR /go/src/github.com/DripEmail/custom-dynamic-weather

RUN apk update && \
    apk add --no-cache git

ENV GO111MODULE on
ENV CGO_ENABLED 0

COPY go.mod go.sum ./
RUN go list -e $(go list -m all)

COPY . .

RUN go build -a -ldflags='-extldflags "-static" -s -w' -o /custom-dynamic-weather

FROM alpine:3.8

RUN apk update && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates

COPY --from=builder custom-dynamic-weather .

EXPOSE 8080

ENTRYPOINT ["./custom-dynamic-weather"]
