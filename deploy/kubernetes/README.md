# Deployment on Kubernetes

For quickly setup Thermo-center on your k8s arm cluster, you will need some minor preparations.

- you will have to prepare nodes with SPI device(s).
- you will need a postgres database

## SPI device/node setup

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
# cat /etc/rc.local
#!/bin/sh

echo 2 > /sys/class/gpio/export

exit 0
```

Then, we must tell this device is available to k8s. For this a device plugin is used. A generic is available at https://github.com/rkojedzinszky/k8s-generic-device-plugin.

Its configuration is easy, assuming you have the spi device at `/dev/spidev0.0` and gpio at `/sys/class/gpio/gpio2`, the following config can be used:
```
# cat /etc/kubernetes/cc1101.yaml
---
name: hardware/cc1101
sets:
- id: cc1101-0
  spec:
    devices:
    - hostpath: /dev/spidev0.0
      containerpath: /dev/spidev0.0
      permissions: "rw"
    mounts:
    - hostpath: /sys/class/gpio/gpio2
      containerpath: /gpio
```

Then you will have to start the device-plugin daemon. To make it permament, use:
```
# cat /etc/systemd/system/k8s-cc1101-resource.service
[Unit]
Description=Kubernetes CC1101 resource provider
After=kubelet.service

[Service]
ExecStart=/usr/local/bin/k8s-generic-device-plugin /etc/kubernetes/cc1101.yaml

[Install]
WantedBy=multi-user.target
```

## Postgres

For Thermo-center to work, an empty postgres database will be needed, the first run of the api component will initialize that.

## Deploy

- To deploy Thermo-center, you will need to have a hostname on which you want to reach its UI. That must be edited in `02-environment.yaml` and in `20-ingress.yaml`. You will have to fill DB parameters in `03-secret.yaml`. Then, just load the files into k8s with:
```
# kubectl -n thermo-center create -f 00-services.yaml -f 02-environment.yaml -f 03-secret.yaml -f 10-deploy.yaml -f 20-ingress.yaml
```

For the first time, a superuser must be created in Django. This is done by executing a django command in the api pod. First, find out the pod's name:
```
# kubectl -n thermo-center get pod -l app=thermo-center-api
NAME                                 READY   STATUS    RESTARTS   AGE
thermo-center-api-77d858c849-svc4j   1/1     Running   0          4m5s
```

Then we use the command:
```
# $ kubectl -n thermo-center exec -it thermo-center-api-77d858c849-svc4j python manage.py createsuperuser
Username (leave blank to use 'thermo'): admin
Email address:
Password:
Password (again):
Superuser created successfully.
```

Then you can log into the UI at `http://<hostname>/tc/`

Django admin can be accessed at `http://<hostname>/tc/admin/`
