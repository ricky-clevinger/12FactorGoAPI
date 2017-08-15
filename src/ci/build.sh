#!/bin/sh
set -e -u -x
cd go-library-source
ls -a
cd src
ls -a
go build main.go
