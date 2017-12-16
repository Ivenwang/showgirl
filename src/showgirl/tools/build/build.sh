#!/usr/bin/env bash


profile=""
while getopts "p:" arg
do
    case $arg in
        p)
            profile=$OPTARG
            ;;
        ?)
            echo $OPTARG "must provide -p value  value can be online,test,dev,demo"
            exit -1
    esac
done

if [ $profile != "dev" ]; then
    #update submodule code
    git submodule update --init --recursive
fi


CURDIR=`dirname $0`
NPATH=`cd $CURDIR/../../../..; pwd`
export GOPATH=$NPATH
echo "GOPATH=$GOPATH"
cd $CURDIR/../..

#
# go 1.9 版本编译 vendor/golang.org/x/net/context/go19.go 类型重复定义
# 解决方案：1.9 Ubuntu 版本需要删除 go19.go 
#
echo $(go version)
echo "Start building!"
rm -rf showgirl
go build

#check exit status
if [ $? != 0 ]; then
    exit -1
fi


if [ ! -f ./showgirl ]; then
  echo Build Failed!
  exit -1
else
  echo "Building done!"
fi
output=output
rm -rf output/*

mkdir -p $output/bin
mkdir -p $output/log
mkdir -p $output/conf
cp showgirl $output/

echo "filter config for $profile"

OLDDIR=`pwd`

rm -f $OLDDIR/conf/app.conf
cd tools/build
go run filter_config.go $OLDDIR/conf/app.conf.tpl  $OLDDIR/conf/"$profile".filter $OLDDIR/conf/app.conf 

#check exit status
if [ $? != 0 ]; then
    exit -1
fi

echo "created app.conf"

cd $OLDDIR

cp -r conf/app.conf $output/conf/
cp -r static $output/
cp -r views $output/

cd $output
tar zcf showgirl_release.tgz --exclude showgirl_release.tgz .

echo "build success"
