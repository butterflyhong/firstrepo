#!/usr/bin/env bash

echo Setting GOPATH ...
CURDIR=`pwd`
DIR="$CURDIR""/go"
echo $DIR
if [ "$DIR" != "$GOPATH" ]
then
    export OLDGOPATH="$GOPATH"
    export GOPATH="$DIR"
    echo "Set GOPATH=${GOPATH}"
fi

echo Finished!

