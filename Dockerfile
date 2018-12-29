FROM python:3.6-alpine AS app-build
MAINTAINER Richard Kojedzinszky <krichy@nmdps.net>

ENV APP_USER=thermo APP_HOME=/opt/thermo-center

RUN mkdir -p $APP_HOME && \
	adduser -D -H -h $APP_HOME $APP_USER

ADD . $APP_HOME/

WORKDIR $APP_HOME

RUN apk add --no-cache -t .build-deps gcc make libffi-dev postgresql-dev git libc-dev linux-headers && \
	pip install -U pip uwsgi -r requirements.txt && \
	apk del .build-deps

RUN apk add --no-cache libpq

RUN python manage.py collectstatic --no-input

RUN python manage.py gen_canjs_models

RUN rm -rf /root/.cache

FROM node:lts-alpine AS frontend

ADD www /work

WORKDIR /work

RUN npm install

COPY --from=app-build /opt/thermo-center/www/models/g/ /work/models/g/

RUN sh build.sh

FROM app-build

RUN apk add --no-cache nginx

RUN mkdir -p /var/www/html/tc/dist/ /var/www/html/tc/icons/ /var/www/html/tc/static/

COPY --from=frontend /work/index.html /var/www/html/tc
COPY --from=frontend /work/dist /var/www/html/tc/dist
COPY --from=frontend /work/icons /var/www/html/tc/icons
COPY --from=app-build /opt/thermo-center/www/static /var/www/html/tc/static

RUN apk add --no-cache supervisor

ADD docker-assets /

ENV APPDAEMON_SOCKET /tmp/appdaemon.sock

EXPOSE 80 8080 8081

ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
