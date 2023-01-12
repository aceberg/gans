FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/gans/ && CGO_ENABLED=0 go build -o /gans .


FROM alpine

WORKDIR /app

RUN apk add --no-cache git openssh-client ca-certificates bash \
    && mkdir -p /data/gans \
    && mkdir -p /root/.ssh

COPY --from=builder /gans /app/
COPY config /root/.ssh/

ENTRYPOINT ["./gans"]