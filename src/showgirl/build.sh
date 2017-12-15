#!/bin/bash
#当增加新的依赖库后,需要执行godep save,将生成的vendor目录业提交到git仓库,
#因为其他同学更新代码后编译时,优先查找vendor目录再查找GOPATH下资源,即可以
#达到无需再次go get新依赖库到本地GOPATH目录下,提高编译效率

CURDIR=`dirname $0`
NPATH=`cd $CURDIR/../..; pwd`
cd $CURDIR

##############################
cd tools/build
sh build_dev.sh
exit
##############################

