# builder
FROM golang:1.11.0-stretch as builder
WORKDIR /gin-gonic
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o app

# build
FROM bitnami/minideb
WORKDIR /gin-gonic

ENV APP_HTTPENDPOINT=:31001
ENV APP_DEV=true
ENV APP_LOG_LEVEL=info
ENV ENABLE_GET=true
ENV ENABLE_PUT=false

# Complete settings, refer to config.go

WORKDIR /root/

COPY --from=builder /gin-gonic/app .
#COPY --from=builder /output/docs ./docs

EXPOSE 8888
CMD ["./app"]
