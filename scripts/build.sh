#!/bin/bash -xe
# (c) 2020 SamCaldwell See LICENSE.txt
# This is the script for running the build-test cycle.
# It is used by azure-pipelines.yml.

echo "Current Working Directory: $(pwd)"

git clone git@github.com:sam-caldwell/adrestia-assertions.git
git checkout development
cd adrestia-assertions

echo "Export GOPATH"
export GOPATH=$(pwd)/adrestia-assertions
echo "--> GOPATH: $GOPATH"

echo "Run Go test"
go test -v
echo "done"
