# thermo-center
DIY thermo solution center component

# Installation

## First, create a python3 virtualenv
```bash
$ VENV=~/.venv
$ virtualenv -p /usr/bin/python3 $VENV
```

## Activate it
```bash
$ . $VENV/bin/activate
```

## Install required python packages
```bash
$ pip install -U -r requirements.txt
```

## Configure the application, postgresql database
```bash
$ umask $(umask && umask 077 && cp local_settings.py.sample local_settings.py)
```

## Do database initialization, populate static files
```bash
$ python manage.py migrate
```

## Add an (admin) user to the system
```bash
$ python manage.py createsuperuser
```

## Additional settings
Edit local_settings.py for other local needs.

### SPI and GPIO setup

Probably you will need to grant permissions to access the devices under /dev, so first create a group for that purpose, add the thermo user to that group, and create the necessary udev rule files:
```bash
# groupadd -r spidev
# adduser thermo spidev
# cat > /etc/udev/rules.d/51.spidev.rules <<EOF
SUBSYSTEM=="spidev", ACTION=="add", GROUP="spidev", MODE="0660"
EOF
```

You will have to select the GPIO for the interrupt-driven packet receiving engine to work. It is up to you which GPIO you choose on your PI, you should connect that to the GDO0 pin of the CC1101 module.
Then you may use the script to expose that GPIO to userspace, with:
```bash
# ~thermo/thermo-center/bin/export-gpio.sh 272 spidev in # Banana PI M1+ GPIO
```

And configure this GPIO in local_settings.py.

Using the attached [adapter](./simple-adapter/) , with one CC1101 module one may use the following table:

| Board | GPIO | pin name |
|-------|-----:|----------|
| Banana PI M1+ | 272 | PI16 |
| Banana PI M2 | 200 | PG8 |
| Banana PI M2+ (H3) | 2 | PA2 |
| Orange PI Plus 2E (H3) | 2 | PA2 |

(Others to be listed here...)

## Frontend setup (www)
You will need an nginx (or other) web server and proxy which can proxy requests to django via uwsgi. The default installation assumes that the application will be reached in the /tc/ URI, which can be set via WWW_ROOT in local_settings.py.
With the default installation, the corresponding nginx configuration looks like this:
```
server {
        listen 80 default_server ;
        listen [::]:80 default_server ;

        root /var/www/html;

        location /tc/ {
                alias /home/thermo/thermo-center/www/;
        }
        location /tc/api/ {
                uwsgi_pass 127.0.0.1:8080;
                include uwsgi_params;
        }
        # Below is only needed if you plan to use Django admin site
        location /tc/admin/ {
                uwsgi_pass 127.0.0.1:8080;
                include uwsgi_params;
        }
}
```

If you plan to deploy it in the root of your domain, set WWW_ROOT to '', and adjust nginx config appropriately.

After, populate django admin static files:
```bash
$ python manage.py collectstatic --noinput --clear
```

And build the web UI files. You will need nodejs and npm to be installed for this to work.
```bash
$ cd www && npm install && sh build.sh
```

## Start the application (the receiver daemon & uwsgi application)
```bash
$ ~/thermo-center/bin/start-daemons.sh
```

If everything is set up correctly, you may open your browser at http://<domain or ip>/tc/ and using the user/password you will see an empty page with 3 buttons on it. Thats all!

# Configuring sensors
TBD
