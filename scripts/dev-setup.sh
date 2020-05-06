#!/bin/bash

export GO111MODULE="on"
go get -u github.com/git-hooks/git-hooks

git hooks install
git hooks list
