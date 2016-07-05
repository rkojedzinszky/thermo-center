#!/bin/sh

ROOT="$(dirname "$(dirname "$(readlink -f "$0")")")"
cd "$ROOT"

export LANG=en_US.UTF-8
export PYTHONHASHSEED=random

~/graphite/bin/carbon-cache.py start

rm -f receiver.sock
python manage.py receiver -d
uwsgi --uwsgi-socket 127.0.0.1:8080 --wsgi-file=application/wsgi.py --master --workers 2 --threads 8 --daemonize=/dev/null
