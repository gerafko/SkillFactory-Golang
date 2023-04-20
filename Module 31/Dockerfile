FROM golang:1.16-buster AS builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 go build  -mod=vendor -o bin/service ./cmd/server/server.go

FROM debian:stable-slim
RUN apt-get update

COPY --from=builder /app/bin/service /service

CMD ["/service"]