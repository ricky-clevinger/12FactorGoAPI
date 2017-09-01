#!/bin/sh
set -e -u -x
cd go-library-source
ls -a

export GOPATH=$PWD
export PATH=$PATH:$GOPATH

cd flyway
flyway validate -configFile=flyway.conf
cd ..

cd src/12FactorGoAPI
ls -a
go build main.go
