# Thermo-center dockerized setup

## Overview

The built docker image has some pre-defined behavior:
* The exported API will be available under /tc/ URI (WWW_ROOT set to 'tc/')

It is possible to configure the application through environment variables only. Then the image assumes the following:
* SPI is accessed at bus 0 / CS0 (/dev/spidev0.0)
* GPIO for interrupt is mounted at /gpio

The available environment variables are as:

| Variable name | Purpose | default value |
|---------------|---------|---------------|
| DBNAME | Postgresql database name | thermo-center |
| DBHOST | Postgresql database host | postgres |
| DBUSER | Postgresql database user | thermo-center |
| DBPASSWORD | Postgresql database password | thermo-center |
| SECRET_KEY | Django SECRET_KEY | change-this |
| DEBUG | Run Django in DEBUG mode | '' |
| ALLOWED_HOSTS | Django ALLOWED_HOSTS, comma separated list | '' |
| CARBON_PICKLE_RECEIVER_HOST | Carbon Pickle receiver host | carbon-cache |
| CARBON_PICKLE_RECEIVER_PORT | Carbon Pickle receiver port | 2004 |
| APPDAEMON_SOCKET | Application daemon control socket (will be deprecated/refactored) | /tmp/appdaemon.sock |
| MQTT_HOST | MQTT host for sensor updates | mqtt |
| MQTT_PORT | MQTT port for sensor updates | 1883 |

Thus the minimal needed docker command is something like this:
```bash
$ docker run \
	-e DBNAME=thermo-docker \
	-e DBHOST=postgres \
	-e DBUSER=thermo-center \
	-e DBPASSWORD=thermo-center \
	-e SECRET_KEY=secret_key \
	-e ALLOWED_HOSTS=thermo-center \
	-e CARBON_PICKLE_RECEIVER_HOST=carbon-cache \
	-e MQTT_HOST=mqtt \
	--device /dev/spidev0.0:/dev/spidev0.0 \
	-v /sys/class/gpio/gpio272:/gpio \
	-p 80:80 \
	--name center \
	-it --rm \
	rkojedzinszky/thermo-center
```

## Installation
For the very first time and each time a new version is to be installed (new image pulled) one should migrate the database. While the new image is running, on another console issue:
```bash
$ docker exec -it center python manage.py migrate
```

Also for the first time one should create a superuser for the application with:
```bash
$ docker exec -it center python manage.py createsuperuser
```

## Timezone

For the heatcontrol feature to work correctly, timezone should be set. The simplest way is to mount the timezone database into the container at /etc/localtime. For example,
    specify the additional arguments to docker:
```bash
$ docker run ... \
	-v /usr/share/zoneinfo/Europe/Budapest:/etc/localtime:ro
	...
```

## Exported ports

The image exports port 80, on which the whole application is available, under /tc/. So after a successful installation one should open http://thermo-center/tc/ in a browser, and it should work.

The image also exports the uwsgi application and the websocket port directly. The uwsgi application listens on port 8080 and the websocket server listens on 8081. If one would like a front-end nginx proxy for this,
one should start the image like this:
```bash
$ docker run \
	-e DBNAME=thermo-docker \
	-e DBHOST=postgres \
	-e DBUSER=thermo-center \
	-e DBPASSWORD=thermo-center \
	-e SECRET_KEY=secret_key \
	-e ALLOWED_HOSTS=thermo-center \
	-e CARBON_PICKLE_RECEIVER_HOST=carbon-cache \
	-e MQTT_HOST=mqtt \
	--device /dev/spidev0.0:/dev/spidev0.0 \
	-v /sys/class/gpio/gpio272:/gpio \
	-p 8080:8080 \
	-p 8081:8081 \
	-p 8082:80 \
	--name center \
	-it --rm \
	rkojedzinszky/thermo-center
```
and use the following nginx configuration snippet:
```
server {
	...

	location /tc/ {
		proxy_pass http://127.0.0.1:8082/tc/;
		proxy_http_version 1.1;
	}
	location /tc/api/ {
		uwsgi_pass 127.0.0.1:8080;
		include uwsgi_params;
	}
	location /tc/admin/ {
		uwsgi_pass 127.0.0.1:8080;
		include uwsgi_params;
	}
	location /tc/ws/ {
		proxy_pass http://127.0.0.1:8081;
		include proxy_params;
		proxy_http_version 1.1;
		proxy_set_header Upgrade $http_upgrade;
		proxy_set_header Connection "upgrade";
	}

	...
}
```
