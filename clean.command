#!/usr/bin/env bash

here="$(dirname "$BASH_SOURCE")"
cd $here
docker rmi kcq/golang-app .
docker rmi kcq/golang-app.slim .