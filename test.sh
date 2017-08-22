#!/bin/bash
set -e

export PATH=$PATH:$PWD
export GOPATH=$PWD/gopath

cwd=$(pwd)

cd ${GOPATH}/src/github.com/EngineerBetter/yml2env

# This env var is required! (and we don't know about set -u)
go test -args -fixtures $FIXTURE_LOCATION -junit "$cwd/test-report/junit-$(date +%s).xml"
