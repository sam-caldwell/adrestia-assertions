#!/bin/bash -e

echo "GOPATH: $GOPATH"

echo "Install GoDoc"
go get golang.org/x/tools/cmd/godoc

echo "Install yamllint"
brew install yamllint

echo "Install golint"
go get -u golang.org/x/lint/golint || {
    echo "HOME: $HOME"
    ls -la $HOME
    exit 1
}
echo "$0 done"