#!/bin/bash

function genDocHtml {
  MDFileName=$1
  PbDirName=$2
  OnlyPublic=$3
  #step1 generate markdown
  cd $NPATH
  rm -rf $MDFileName.md
  ./md_generator $PbDirName $MDFileName.md $OnlyPublic

  #step2 convert markdown to html by tocmd
  cd tocmd.npm-master
  rm -rf preview
  node bin/tocmd.js -f ../$MDFileName.md -o $MDFileName.html
  echo cp preview/$MDFileName.html $PbDirName/
  cp preview/$MDFileName.html $PbDirName/
}

CURDIR=`dirname $0`
NPATH=`cd $CURDIR; pwd`
cd $NPATH

PROTO=/opt/soft/jenkins_data/jobs/tiantian-golang/workspace/proto
if [ ! x$1 = x ]; then
  PROTO=$1
fi
echo PROTO=$PROTO

# gen api.html
genDocHtml api $PROTO 1

# gen all.html
genDocHtml all $PROTO 0

echo "Execute Done"