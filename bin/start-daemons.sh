#!/bin/sh

export LANG=en_US.UTF-8
export PYTHONHASHSEED=random

~/graphite/bin/carbon-cache.py start
