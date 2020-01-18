FROM golang:1.13-alpine AS go-build

MAINTAINER Richard Kojedzinszky <richard@kojedz.in>

RUN apk --no-cache add git binutils

WORKDIR /go/src/thermo-center-ws

ADD ws .

RUN go get -d && \
    CGO_ENABLED=0 go build && \
    strip -s thermo-center-ws

FROM scratch

COPY --from=go-build /go/src/thermo-center-ws/thermo-center-ws /

USER 65535

EXPOSE 8081

CMD ["/thermo-center-ws"]
