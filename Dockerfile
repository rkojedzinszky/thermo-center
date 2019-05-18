FROM alpine:3.9 AS app-build
MAINTAINER Richard Kojedzinszky <richard@kojedz.in>

# Install python3
RUN apk add --no-cache python3 && \
    ln -sf python3 /usr/bin/python && ln -sf pip3 /usr/bin/pip && \
    pip install -U pip

ENV APP_USER=thermo APP_HOME=/opt/thermo-center

RUN mkdir -p $APP_HOME && \
	adduser -S -D -H -h $APP_HOME $APP_USER

ADD requirements.txt manage.py $APP_HOME/
ADD application $APP_HOME/application
ADD center $APP_HOME/center
ADD heatcontrol $APP_HOME/heatcontrol
ADD nauth $APP_HOME/nauth

WORKDIR $APP_HOME

RUN apk add --no-cache py3-django py3-psycopg2 py3-ipaddress py3-crypto uwsgi-python3 && \
    apk add --no-cache -t .build-deps gcc make libffi-dev git libc-dev linux-headers python3-dev && \
    pip install -r requirements.txt && \
    apk del .build-deps && \
    rm -rf /root/.cache

RUN python manage.py collectstatic --no-input && \
    mkdir -p www/models/g && \
    python manage.py gen_canjs_models

FROM node:lts-alpine AS frontend

ADD www /work

WORKDIR /work

COPY --from=app-build /opt/thermo-center/www/models/g/ /work/models/g/

RUN yarn && sh build.sh && rm -rf node_modules

FROM app-build

RUN apk add --no-cache nginx tzdata && \
    pip install supervisor && \
    rm -rf /root/.cache

RUN mkdir -p /var/www/html/tc/dist/ /var/www/html/tc/icons/ /var/www/html/tc/static/

COPY --from=frontend /work/index.html /var/www/html/tc
COPY --from=frontend /work/dist /var/www/html/tc/dist
COPY --from=frontend /work/icons /var/www/html/tc/icons
COPY --from=app-build /opt/thermo-center/www/static /var/www/html/tc/static

ADD docker-assets /

ENV APPDAEMON_SOCKET /tmp/appdaemon.sock

EXPOSE 80 8080 8081

ENTRYPOINT ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
