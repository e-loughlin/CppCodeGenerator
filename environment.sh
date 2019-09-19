#!/bin/bashi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
export GOPATH=$DIR
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
