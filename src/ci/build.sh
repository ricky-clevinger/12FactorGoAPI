#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
export GOPATH=/usr/local/go
export PATH=$PATH:$GOPATH
cd src
ls -a
go build main.go
