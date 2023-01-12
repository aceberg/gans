FROM golang:alpine AS builder

RUN apk add build-base
COPY cmd /src/cmd
COPY internal /src/internal
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src/cmd/gans/ && CGO_ENABLED=0 go build -o /gans .


FROM alpine

WORKDIR /app

RUN apk add --no-cache git ansible openssh-client ca-certificates bash sshpass nano \
    && mkdir -p /data/gans 

COPY --from=builder /gans /app/
COPY config/ssh_config /etc/ssh/ssh_config

ENTRYPOINT ["./gans"]