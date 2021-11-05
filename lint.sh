#!/usr/bin/env bash

set -eu

mkdir -p /go/src/github.com/EngineerBetter
cp -r cli-code /go/src/github.com/EngineerBetter/yml2env
cd /go/src/github.com/EngineerBetter/yml2env
gometalinter --disable-all --enable=ineffassign --enable=deadcode
