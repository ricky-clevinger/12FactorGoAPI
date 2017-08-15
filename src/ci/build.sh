#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
cd src
export GOPATH=$PWD
export PATH=$PATH:$GOPATH
ls -a
go build main.go
