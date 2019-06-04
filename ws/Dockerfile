FROM rkojedzinszky/alpine-python-grpcio

MAINTAINER Richard Kojedzinszky <richard@kojedz.in>

ENV APP_HOME=/opt/thermo-center-ws

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

ADD ws ws

RUN pip install -r ws/requirements.txt && \
    rm -rf /root/.cache

USER nobody

EXPOSE 8081

CMD ["python", "-m", "ws.main"]
