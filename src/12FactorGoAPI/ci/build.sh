#!/bin/sh
set -e -u -x
cd go-library-source
ls -a

git clone git://github.com/pliniocsfernandes/flyway-command-line.git /opt/flyway
ln -s /opt/flyway/flyway.sh /usr/bin/flyway

export GOPATH=$PWD
export PATH=$PATH:$GOPATH

flyway validate

cd src/12FactorGoAPI
ls -a
go build main.go
