FROM scratch

LABEL org.opencontainers.image.authors "Richard Kojedzinszky <richard@kojedz.in>"
LABEL org.opencontainers.image.source https://github.com/rkojedzinszky/thermo-center

COPY receiver /

EXPOSE 8079

USER 8079

CMD ["/receiver"]
