#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
export GOROOT=/usr/local/go
export GOPATH=$PWD
export PATH=$PATH:$GOPATH:$GOROOT/bin
cd src
ls -a
go build main.go
