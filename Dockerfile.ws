FROM scratch

LABEL org.opencontainers.image.authors "Richard Kojedzinszky <richard@kojedz.in>"
LABEL org.opencontainers.image.source https://github.com/rkojedzinszky/thermo-center

ADD bin/ws /

EXPOSE 8081

USER 65535

CMD ["/ws"]
