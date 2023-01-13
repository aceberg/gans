FROM golang:alpine AS builder

RUN apk add build-base
COPY cmd /src/cmd
COPY internal /src/internal
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src/cmd/gans/ && CGO_ENABLED=0 go build -o /gans .


FROM alpine

WORKDIR /app

RUN apk add --no-cache tzdata git ansible openssh-client sshpass \
    ca-certificates bash nano && \
    mkdir -p /data/gans 

COPY --from=builder /gans /app/

ENTRYPOINT ["./gans"]