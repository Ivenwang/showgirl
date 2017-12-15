#!/bin/bash
CURDIR=`dirname $0`
NPATH=`cd $CURDIR/../..; pwd`
CURDIR=$NPATH/src/showgirl
export GOPATH=$NPATH
echo "GOPATH=$GOPATH"
cd $CURDIR
OS=`uname`
if [ "x"$OS = "x"'MINGW64_NT-6.1' ]; then
  OS='Windows_NT'
fi

if [ "x"$OS = "x"'MINGW32_NT-6.1-WOW' ]; then
  OS='Windows_NT'
fi


if [ "x"$OS = "x"'Windows_NT' ]; then
  echo "cd $NPATH/proto && $NPATH/src/showgirl/tools/$OS/protoc.exe --plugin=$CURDIR/tools/$OS/protoc-gen-go.exe --go_out=$CURDIR/client *.proto"
  cd $NPATH/proto && $NPATH/src/showgirl/tools/$OS/protoc.exe --plugin=$CURDIR/tools/$OS/protoc-gen-go.exe --go_out=$CURDIR/client *.proto
else
  # LD_LIBRARY_PATH 程序加载运行期间查找动态链接库时指定除了系统默认路径之外的其他路径
  export LD_LIBRARY_PATH=$NPATH/src/showgirl/tools/$OS/.libs:$LD_LIBRARY_PATH
  echo "cd $NPATH/proto && $NPATH/src/showgirl/tools/$OS/protoc --plugin=$CURDIR/tools/$OS/protoc-gen-go --go_out=$CURDIR/client *.proto"
  cd $NPATH/proto && $NPATH/src/showgirl/tools/$OS/protoc --plugin=$CURDIR/tools/$OS/protoc-gen-go --go_out=$CURDIR/client *.proto
fi

cd $CURDIR
#build autocode if needed
if [ ! -f tools/$OS/autocode ]; then
  sh tools/code_generator/build.sh
  cp tools/code_generator/code_generator tools/$OS/autocode
  if [ ! -f tools/$OS/autocode ]; then
    echo "code_generator build failed!"
    exit
  fi
fi

#generate auto code if needed
cd $CURDIR
cd tools;./$OS/autocode;cd -
echo "code generate done!"

