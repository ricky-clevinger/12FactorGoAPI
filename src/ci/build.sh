#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
export PATH=$PATH:$GOPATH
cd src
export GOPATH=$PWD
ls -a
go build main.go
