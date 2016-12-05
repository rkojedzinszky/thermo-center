#!/bin/sh

set -e
test -d dist && find dist -mindepth 1 -maxdepth 1 -print0 | xargs -r0 rm -rf
nodejs build.js
find dist/ -type f -regex '.*.\(js\|css\)' -print0 | xargs -t -r0 gzip -k9
