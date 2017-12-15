#!/bin/bash
# 1. You should install tocmd.
#    gem install tocmd
# 2. Run this script in any dir

CURDIR=`dirname $0`
NPATH=`cd $CURDIR; pwd`
cd $NPATH/../../../../proto

MDFileName=./api.md

#step1 generate markdown
$NPATH/md_generator . $MDFileName 1

#step2 convert markdown to html by tocmd
tocmd -f $MDFileName

#step3 add scroll bar for side directory tree
perl -pi -e "s#class=\"ztree\"\ style=\'width\:100\%#class=\"ztree\"\ style=\'width\:100\%\;height\:100\%\;overflow-y\:auto\;#gi" preview/api.html

MDFileName=./all.md
#step1 generate markdown
$NPATH/md_generator . $MDFileName 0

#step2 convert markdown to html by tocmd
tocmd -f $MDFileName

#step3 add scroll bar for side directory tree
perl -pi -e "s#class=\"ztree\"\ style=\'width\:100\%#class=\"ztree\"\ style=\'width\:100\%\;height\:100\%\;overflow-y\:auto\;#gi" preview/all.html

