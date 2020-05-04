#!/bin/bash -xe
# (c) 2020 SamCaldwell See LICENSE.txt
# This is the script for running the build-test cycle.
# It is used by azure-pipelines.yml.

ls -lah

export GOPATH=/home/vsts/work/1/s/

go test -v
