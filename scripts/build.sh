#!/bin/bash -xe
# (c) 2020 SamCaldwell See LICENSE.txt
# This is the script for running the build-test cycle.
# It is used by azure-pipelines.yml.

echo "Current Working Directory: $(pwd)"

export GOPATH=$(pwd)/git
echo "Export GOPATH"
echo "--> GOPATH: $GOPATH"

mkdir -p git/sam-caldwell
cd git/sam-caldwell || exit 1

git clone git@github.com:sam-caldwell/adrestia-assertions.git

cd adrestia-assertions || exit 1
git checkout development
git checkout $(git log -n1 | head -n1 | awk '{print $2}') -b build

echo "Run Go test"
go test -a -v -test.benchmem -test.parallel 2
echo "done"
