#!/bin/bash -xe
# (c) 2020 SamCaldwell See LICENSE.txt
# This is the script for running the build-test cycle.
# It is used by azure-pipelines.yml.

ls -lah
cd /home/vsts/work/1/s/github.com/sam-caldwell/adrestia-assertions/

export GOPATH=/home/vsts/work/1/s/
ls -la $GOPATH

go test -v
