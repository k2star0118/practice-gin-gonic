# builder
FROM innobead/go-gradle-docker:1.11.1 as builder
COPY . /output

# build
FROM bitnami/minideb
WORKDIR /project

ENV APP_HTTPENDPOINT=:31001
ENV APP_DEV=true
ENV APP_LOG_LEVEL=info
# Complete settings, refer to config.go

COPY --from=builder /output/.gogradle/gin-gonic .
COPY --from=builder /output/docs ./docs

CMD ["/project/gin-gonic"]
