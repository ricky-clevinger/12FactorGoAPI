#!/bin/sh
set -e -u -x
cd go-library-source
ls -a

export GOPATH=$PWD
export PATH=$PATH:$GOPATH
cd src/12FactorGoAPI

flyway validate -configFile=flyway.conf

ls -a
go build main.go
