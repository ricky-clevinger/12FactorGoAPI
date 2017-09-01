#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
export GOPATH=$PWD
export PATH=$PATH:$GOPATH
cd src/12FactorGoAPI
ln -s /flyway/flyway.sh /usr/local/bin/flyway
ls -a
go build main.go
