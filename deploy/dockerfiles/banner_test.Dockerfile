FROM golang:1.22.1-alpine AS builder

WORKDIR /build

COPY go.mod .

COPY . .

RUN go build -o banner_service ./cmd/service/main.go

FROM ubuntu:20.04

ENV APP_DIR /build

WORKDIR $APP_DIR

RUN apt-get update \
    && groupadd -r web \
    && useradd -d $APP_DIR -r -g web web \
    && chown web:web -R $APP_DIR \
    && apt-get install -y netcat-traditional \
    && apt-get install -y acl

COPY --from=builder /build/banner_service $APP_DIR/banner_service
COPY --from=builder /build/deploy/scripts/test-banner-service-start.sh $APP_DIR/test-banner-service-start.sh
COPY --from=builder /build/deploy/env/.env.test $APP_DIR//deploy/env/.env.test

RUN setfacl -R -m u:web:rwx $APP_DIR/test-banner-service-start.sh

USER web

ENTRYPOINT ["bash", "test-banner-service-start.sh"]