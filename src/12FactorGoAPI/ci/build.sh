#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
export GOPATH=$PWD
export PATH=$PATH:$GOPATH
cd src/12FactorGoAPI

sudo git clone git://github.com/pliniocsfernandes/flyway-command-line.git /opt/flyway
sudo ln -s /opt/flyway/flyway.sh /usr/bin/flyway

ls -a
go build main.go
