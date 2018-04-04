#!/bin/sh

ROOT="$(dirname "$(dirname "$(readlink -f "$0")")")"
cd "$ROOT"

export LANG=en_US.UTF-8

rm -f receiver.sock
python manage.py receiver -d
uwsgi --ini uwsgi.ini
