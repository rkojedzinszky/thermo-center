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
ADD nauth nauth
ADD aggregator/pid.py aggregator/pid.py

RUN apk add --no-cache tzdata py3-psycopg2 && \
    pip install -r requirements.txt && \
    rm -rf /root/.cache

### API
FROM common AS api

RUN apk add --no-cache uwsgi-python3 uwsgi-cheaper_busyness

EXPOSE 8080 8082

USER $APP_USER

CMD ["uwsgi", \
    "--need-plugin=python3", \
    "--need-plugin=cheaper_busyness", \
    "--uwsgi-socket=:8080", \
    "--http-socket=:8082", \
    "--wsgi-file=application/wsgi.py", \
    "--master", \
    "--die-on-term", \
    "--workers=2", \
    "--threads=4", \
    "--cheaper=1", \
    "--cheaper-algo=busyness", \
    "--cheaper-overload=10", \
    "--thunder-lock"]

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
