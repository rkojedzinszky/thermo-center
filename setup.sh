#!/bin/sh

pip install -U -r requirements.txt

# setup local_settings.py
if [ ! -f local_settings.py ]; then
    echo "** Generating local_settings.py"
    SECRET_KEY=$(python -c 'import random; print "".join(random.choice("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^*(-_=+)") for i in range(50))')
    CACHE_DIR=$(python -c 'import random; print "".join(random.choice("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") for i in range(12))')
    _u=$(umask)
    umask 027
    sed \
	    -e "s/@SECRET_KEY@/$SECRET_KEY/g" \
	    -e "s#@CACHE_DIR@#$CACHE_DIR#g" \
	    local_settings.py.sample > local_settings.py
    umask $_u
fi

python manage.py migrate
