#!/bin/bash
CURDIR=`dirname $0`
NPATH=`cd $CURDIR/../../../..; pwd`
export GOPATH=$NPATH
echo "GOPATH=$GOPATH"
cd $CURDIR
go get github.com/astaxie/beego
echo "Start building md_generator!"
rm md_generator
go build
ls -l md_generator
if [  $? -ne 0 ]; then
  echo "Build code generator failed!"
else
  echo "Build code generator done!"
fi
