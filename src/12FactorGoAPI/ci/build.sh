#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
export GOPATH=$PWD
export PATH=$PATH:$GOPATH;go-library-source/src/12FactorGoAPI
cd src/12FactorGoAPI

ls -a
go build main.go
