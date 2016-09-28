#!/bin/sh

set -e
test -d dist && find dist -type f -name '*.gz' -delete
nodejs build.js
find dist/ -type f -regex '.*.\(js\|css\)' -print0 | xargs -t -r0 gzip -k9
