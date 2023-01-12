FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/gans/ && CGO_ENABLED=0 go build -o /gans .


FROM alpine

WORKDIR /app

RUN apk add --no-cache git ansible openssh-client ca-certificates bash sshpass \
    && mkdir -p /data/gans 

COPY --from=builder /gans /app/

ENTRYPOINT ["./gans"]