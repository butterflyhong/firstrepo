#!/usr/bin/env bash

echo Setting GOPATH ...
DIR=`pwd`/go
if [ "$DIR" != "$GOPATH" ]
then
    export OLDGOPATH=$GOPATH
    export GOPATH=$DIR
    echo "Set GOPATH=${GOPATH}"
fi

echo Finished!

