FROM golang:1.12 AS build-env

# RUN apk add --no-cache git

RUN adduser --disabled-login --uid 42987 app && mkdir -p /app/ && chown app /app/
USER app

WORKDIR /app/
ADD . /app/

# Disable Cgo
RUN GO_ENABLED=0 go build -o /app/app .

# Run-time:
FROM alpine:3.8

RUN adduser -D -u 42987 app && mkdir -p /app/ && chown app /app/
USER app

WORKDIR /
COPY --from=build-env /app/server.* /
COPY --from=build-env /app/app /

EXPOSE 8080
CMD ["/app"]

