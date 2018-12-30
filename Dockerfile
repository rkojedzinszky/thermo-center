FROM python:3.6-alpine AS app-build
MAINTAINER Richard Kojedzinszky <krichy@nmdps.net>

ENV APP_USER=thermo APP_HOME=/opt/thermo-center

RUN mkdir -p $APP_HOME && \
	adduser -S -D -H -h $APP_HOME $APP_USER

ADD . $APP_HOME/

WORKDIR $APP_HOME

RUN apk add --no-cache -t .build-deps gcc make libffi-dev postgresql-dev git libc-dev linux-headers && \
	pip install -U pip uwsgi -r requirements.txt && \
	apk del .build-deps && \
	rm -rf /root/.cache

RUN apk add --no-cache libpq && \
	python manage.py collectstatic --no-input && \
	python manage.py gen_canjs_models

FROM node:lts-alpine AS frontend

ADD www /work

WORKDIR /work

COPY --from=app-build /opt/thermo-center/www/models/g/ /work/models/g/

RUN npm install && sh build.sh && rm -rf node_modules

FROM app-build

RUN apk add --no-cache nginx supervisor

RUN mkdir -p /var/www/html/tc/dist/ /var/www/html/tc/icons/ /var/www/html/tc/static/

COPY --from=frontend /work/index.html /var/www/html/tc
COPY --from=frontend /work/dist /var/www/html/tc/dist
COPY --from=frontend /work/icons /var/www/html/tc/icons
COPY --from=app-build /opt/thermo-center/www/static /var/www/html/tc/static

ADD docker-assets /

ENV APPDAEMON_SOCKET /tmp/appdaemon.sock

EXPOSE 80 8080 8081

ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
