#!/bin/sh

ROOT="$(dirname "$(dirname "$(readlink -f "$0")")")"
cd "$ROOT"

export LANG=en_US.UTF-8
export PYTHONHASHSEED=random

~/graphite/bin/carbon-cache.py start

rm -f receiver.sock
python manage.py receiver
