#!/bin/sh
set -e -u -x
cd go-library-source
ls -a

git clone git://github.com/pliniocsfernandes/flyway-command-line.git 

export GOPATH=$PWD
export PATH=$PATH:$GOPATH:go-library-source/src/12FactorGoAPI
cd src/12FactorGoAPI

flyway validate -configFile=flyway.conf

ls -a
go build main.go
