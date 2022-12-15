FROM golang:1.18-alpine3.15 as build-stage

RUN apk add --no-cache --update gcc musl-dev git librdkafka-dev

WORKDIR /app

COPY . .

RUN go build -a -tags musl -o fake-data-producer

FROM alpine:3.15

WORKDIR /app

COPY --from=build-stage /app/fake-data-producer /usr/local/bin/app

ENTRYPOINT ["app"]