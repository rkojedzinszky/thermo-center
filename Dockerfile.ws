FROM scratch

ADD bin/ws /

EXPOSE 8081

USER 65535

ENV GOGC=50

CMD ["/ws"]
