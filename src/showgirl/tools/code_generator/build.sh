#!/bin/bash
CURDIR=`dirname $0`
NPATH=`cd $CURDIR/../../../..; pwd`
export GOPATH=$NPATH
echo "GOPATH=$GOPATH"
cd $CURDIR
go get github.com/astaxie/beego
echo "Start building code generator!"
rm code_generator
go build
ls -l code_generator
echo "Build code generator done!"
