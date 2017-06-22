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

## Do database initialization
```bash
$ python manage.py migrate
```

## Additional settings
Edit local_settings.py for other local needs, especially for SPI device access.
