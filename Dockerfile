FROM rkojedzinszky/alpine-python-grpcio AS common
MAINTAINER Richard Kojedzinszky <richard@kojedz.in>

ENV APP_USER=thermo APP_HOME=/opt/thermo-center

RUN mkdir -p $APP_HOME && \
	adduser -D -H -h $APP_HOME $APP_USER

WORKDIR $APP_HOME

ADD lib lib
ADD requirements.txt manage.py ./
ADD application application
ADD center center
ADD heatcontrol heatcontrol
ADD aggregator/pid.py aggregator/pid.py
ADD receiver/api_pb2.py receiver/api_pb2_grpc.py receiver/
ADD configurator/api_pb2.py configurator/

RUN apk add --no-cache tzdata py3-psycopg2 && \
    pip install -r requirements.txt && \
    rm -rf /root/.cache

### API
FROM common AS api

ADD uwsgi.api.ini ./

RUN apk add --no-cache uwsgi-python3 uwsgi-cheaper_busyness

EXPOSE 8080

USER $APP_USER

CMD ["uwsgi", "--ini", "uwsgi.api.ini"]

### APP
FROM common AS grpcserver

ADD configurator configurator
ADD aggregator aggregator

EXPOSE 8079

USER $APP_USER

CMD ["python", "manage.py", "grpcserver", "--configurator", "--aggregator"]

### UI
FROM common AS fe-prepare

RUN python manage.py collectstatic --no-input && \
    mkdir -p www/models/g && \
    python manage.py gen_canjs_models

FROM node:lts-alpine AS fe-build

ADD www /work

WORKDIR /work

COPY --from=fe-prepare /opt/thermo-center/www/models/g/ /work/models/g/

RUN yarn && sh build.sh && rm -rf node_modules

FROM nginx:alpine AS ui

RUN mkdir -p /var/www/html/tc/dist/ /var/www/html/tc/icons/ /var/www/html/tc/static/

ADD www/index.html /var/www/html/tc/
ADD www/icons /var/www/html/tc/icons/
COPY --from=fe-prepare /opt/thermo-center/www/static /var/www/html/tc/static/
COPY --from=fe-build /work/dist /var/www/html/tc/dist

ADD docker-assets-ui /
