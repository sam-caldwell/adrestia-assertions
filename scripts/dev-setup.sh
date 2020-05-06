#!/bin/bash -e
export GO111MODULE="on"
export GOPATH="$GOPATH:$(pwd)/"

echo "GOPATH: $GOPATH"

echo "Install GoDoc"
go get golang.org/x/tools/cmd/godoc

echo "Install yamllint"
brew install yamllint

echo "Install golint"
go get -u golang.org/x/lint/golint

echo "$0 done"