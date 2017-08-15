#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
export GOPATH=go-library-source
export PATH=$PATH:$GOPATH
cd src
ls -a
go build main.go
