FROM rkojedzinszky/alpine-python-grpcio:3.11 AS common
MAINTAINER Richard Kojedzinszky <richard@kojedz.in>

ENV APP_USER=thermo APP_HOME=/opt/thermo-center APP_UID=10101

RUN mkdir -p $APP_HOME && \
	adduser -D -H -h $APP_HOME -u $APP_UID $APP_USER

WORKDIR $APP_HOME

ADD lib lib
ADD requirements.txt manage.py ./
ADD application application
ADD center center
ADD heatcontrol heatcontrol
ADD aggregator/pid.py aggregator/pid.py
ADD receiver/api_pb2.py receiver/api_pb2_grpc.py receiver/
ADD configurator/api_pb2.py configurator/

RUN apk add --no-cache tzdata py3-psycopg2 libmemcached && \
    apk --no-cache add -t .build-deps gcc libc-dev python3-dev zlib-dev libmemcached-dev && \
    pip install -r requirements.txt && \
    apk del .build-deps && \
    rm -rf /root/.cache

### API
FROM common AS api

ADD uwsgi.api.ini ./

RUN apk add --no-cache uwsgi-python3 uwsgi-cheaper_busyness

EXPOSE 8080

USER $APP_UID

CMD ["uwsgi", "--ini", "uwsgi.api.ini"]

### APP
FROM common AS grpcserver

ADD configurator configurator
ADD aggregator aggregator

EXPOSE 8079

USER $APP_UID

CMD ["python", "manage.py", "grpcserver", "--configurator", "--aggregator"]

### UI
FROM common AS fe-prepare

RUN python manage.py collectstatic --no-input && \
    mkdir -p www/models/g && \
    python manage.py gen_canjs_models

FROM node:10-alpine AS fe-build

ADD www /work

WORKDIR /work

COPY --from=fe-prepare /opt/thermo-center/www/models/g/ /work/models/g/

RUN yarn && sh build.sh && rm -rf node_modules

FROM nginx:alpine AS ui

RUN mkdir -p /var/www/html/dist/ /var/www/html/icons/ /var/www/html/static/

ADD www/index.html www/manifest.json /var/www/html/
ADD www/icons /var/www/html/icons/
COPY --from=fe-prepare /opt/thermo-center/www/static /var/www/html/static/
COPY --from=fe-build /work/dist /var/www/html/dist

ADD docker-assets-ui /

# Tune for rootless
RUN sed -r -i \
        -e '/^user/s/^/#/' \
        -e '/^pid/s!.*!pid /tmp/nginx.pid;!' \
	/etc/nginx/nginx.conf && \
	chown 8080:8080 /var/cache/nginx

USER 8080

EXPOSE 8080
