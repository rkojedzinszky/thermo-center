#!/bin/sh

GRAPHITE_ROOT=$HOME/graphite
PYTHON_LIB=$(python -c "from distutils.sysconfig import get_python_lib; print(get_python_lib())")

pip install -U -r requirements.txt

# set up carbon
pip install -U carbon --install-option="--prefix=$GRAPHITE_ROOT" --install-option="--install-lib=$PYTHON_LIB"
(
cd $GRAPHITE_ROOT/conf
if ! test -f carbon.conf ; then
	sed -r \
		-e 's/^MAX_UPDATES_PER_SECOND.*/MAX_UPDATES_PER_SECOND = 10/' \
		-e 's/^LINE_RECEIVER_PORT.*/LINE_RECEIVER_PORT = 0/' \
		-e 's/^PICKLE_RECEIVER_INTERFACE.*/PICKLE_RECEIVER_INTERFACE = 127.0.0.1/' \
		-e 's/^CACHE_QUERY_INTERFACE.*/CACHE_QUERY_INTERFACE = 127.0.0.1/' \
		-e 's/^[#[:space:]]*MAX_UPDATES_PER_SECOND_ON_SHUTDOWN.*/MAX_UPDATES_PER_SECOND_ON_SHUTDOWN = 1000/' \
		carbon.conf.example > carbon.conf
fi

if ! test -f storage-schemas.conf ; then
	sed -r \
		-e '/default_1min_for_1day/,$d' \
		storage-schemas.conf.example > storage-schemas.conf
	cat <<-EOF >>storage-schemas.conf
	[default]
	pattern = .*
	retentions = 5m:15d, 30m:90d, 3h:540d, 1d:12y
	EOF
fi
)

# set up graphite-web
pip install -U graphite-web --install-option="--prefix=$GRAPHITE_ROOT" --install-option="--install-lib=$PYTHON_LIB"

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
	    -e "s#@GRAPHITE_ROOT@#$GRAPHITE_ROOT#g" \
	    local_settings.py.sample > local_settings.py
    umask $_u
fi

python manage.py migrate
