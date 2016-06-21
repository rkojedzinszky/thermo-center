#!/bin/sh

GRAPHITE=$HOME/graphite
PYTHON_LIB=$(python -c 'import sys; print sys.path.pop()')

pip install -U -r requirements.txt

# set up carbon
pip install -U carbon --install-option="--prefix=$GRAPHITE" --install-option="--install-lib=$PYTHON_LIB"
(
cd $GRAPHITE/conf
if ! test -f carbon.conf ; then
	sed -r \
		-e 's/^MAX_UPDATES_PER_SECOND.*/MAX_UPDATES_PER_SECOND = 1/' \
		-e 's/^LINE_RECEIVER_PORT.*/LINE_RECEIVER_PORT = 0/' \
		-e 's/^PICKLE_RECEIVER_INTERFACE.*/PICKLE_RECEIVER_INTERFACE = 127.0.0.1/' \
		-e 's/^CACHE_QUERY_INTERFACE/CACHE_QUERY_INTERFACE = 127.0.0.1/' \
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

## set up graphite-web
#pip install -U graphite-web --install-option="--prefix=$GRAPHITE" --install-option="--install-lib=$PYTHON_LIB"
