#!/bin/sh

# export-gpio <gpio> <group> <dir>

if [ $# -ne 3 ]; then
	echo "Usage: $0 <gpio> <group> <direction>"
	exit 1
fi

cd /sys/class/gpio
echo $1 > export
cd gpio$1
chgrp $2 direction value edge 2>/dev/null
chmod 660 direction value edge 2>/dev/null
echo $3 > direction
