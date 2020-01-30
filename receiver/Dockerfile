FROM rkojedzinszky/alpine-python-grpcio

MAINTAINER Richard Kojedzinszky <richard@kojedz.in>

ENV APP_HOME=/opt/thermo-center-receiver

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

ADD lib lib
ADD configurator/api_pb2.py configurator/api_pb2_grpc.py configurator/
ADD aggregator/api_pb2.py aggregator/api_pb2_grpc.py aggregator/
ADD receiver receiver

RUN apk --no-cache add py3-crypto && \
    apk add --no-cache -t .build-deps python3-dev gcc make libffi-dev libc-dev && \
    pip install -r receiver/requirements.txt && \
    apk del .build-deps && \
    rm -rf /root/.cache

USER 10102

EXPOSE 8079

CMD ["python", "-m", "receiver.main"]
