FROM scratch

ARG TARGETARCH

COPY receiver.${TARGETARCH} /receiver

EXPOSE 8079

USER 8079

CMD ["/receiver"]
