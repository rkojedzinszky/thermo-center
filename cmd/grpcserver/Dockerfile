FROM golang:1.14-alpine AS builder

MAINTAINER Richard Kojedzinszky <richard@kojedz.in>

RUN apk --no-cache add binutils

RUN mkdir -p /go/src/github.com/rkojedzinszky/thermo-center

WORKDIR /go/src/github.com/rkojedzinszky/thermo-center

ADD go.mod go.mod
ADD models models
ADD cmd cmd
ADD configurator configurator
ADD aggregator aggregator

RUN CGO_ENABLED=0 go build ./cmd/grpcserver && \
    strip -s grpcserver

FROM alpine:3.11 AS grpcserver

RUN apk --no-cache add tzdata

COPY --from=builder /go/src/github.com/rkojedzinszky/thermo-center/grpcserver /

EXPOSE 8079

USER 8079

CMD ["/grpcserver"]
