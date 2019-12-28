# Builder
FROM golang:1.13.5-alpine3.11 as builder

RUN apk update && apk upgrade && \
    apk --update add git make build-base

WORKDIR /app

COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /app/rest-server /app
COPY --from=builder /app/grpc-server /app

CMD /app/rest-server