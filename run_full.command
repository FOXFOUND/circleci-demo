#!/usr/bin/env bash

here="$(dirname "$BASH_SOURCE")"
cd $here
docker run --rm -it -p 1300:1300 kcq/golang-app
