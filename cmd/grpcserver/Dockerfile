FROM scratch

ARG TARGETARCH

COPY grpcserver.${TARGETARCH} /grpcserver

EXPOSE 8079

USER 8079

CMD ["/grpcserver"]
