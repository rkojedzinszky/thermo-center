# Deployment with docker compose

For quickly setup Thermo-center with docker-compose, you will need some minor preparations.

- you will have to prepare your SPI device for use
- you will need a postgres database
- you will need to set up a front-end proxy.

## SPI device setup

You must enable spidev on yoar board, on armbian you typically use `armbian-config` and/or edit `/boot/armbianEnv.txt` directly. For example on a H3 based board you would have the following lines:

```
overlay_prefix=sun8i-h3
overlays=spi-spidev
param_spidev_spi_bus=0
```

When your board boots up, you should have an spi device like:
```
# ls -la /dev/spidev0.0
crw------- 1 root root 153, 0 Nov  3  2016 /dev/spidev0.0
```

Also, for receiving interrupts from CC1101 module, we will need a GPIO port. On H3 boards, with the example adapter, you would use the `A2` pin, this can be exported to userspace as:
```
# echo 2 > /sys/class/gpio/export
```

To persist this and grant permissions to rootless containers, the following `rc.local` snippet can be used:

```
# cat /etc/rc.local
#!/bin/sh

echo 2 > /sys/class/gpio/export
chmod 666 /dev/spidev0.0 /sys/class/gpio/gpio2/*

exit 0
```

This two paths are mounted into the `receiver` service. If you have different paths, adjust accordingly.

## Postgres

For Thermo-center to work, an empty postgres database will be needed, the first run of the api component will initialize that.

## Deploy

Edit variables in .env accordingly. Then, for first run, or later for database upgrade, run the following:
```
docker run -it --rm --env-file .env rkojedzinszky/thermo-center-api python manage.py migrate
```

And for the very frist time, a superuser must be created:
```
docker run -it --rm --env-file .env rkojedzinszky/thermo-center-api python manage.py createsuperuser
Username (leave blank to use 'thermo'): admin
Email address:
Password:
Password (again):
Superuser created successfully.
```

Then start the application:
```
# docker-compose up
```

## Front-end proxy

With the sample `nginx.conf.sample`, you can run a reverse proxy on the same host as:
```
# nginx -c $(pwd)/nginx.conf.sample
```

Then you can log into the UI at `http://<hostname>/`

Django admin can be accessed at `http://<hostname>/admin/`

<hostname> is the hostname set in `.env` and points to the running proxy.
