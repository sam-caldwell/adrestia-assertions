#!/bin/bash

export GO111MODULE="on"

echo "Install git-hooks"
go get github.com/git-hooks/git-hooks

echo "Install the git hooks to the local environment."
git hooks install

echo "List current hooks"
git hooks list

echo "Install GoDoc"
go get golang.org/x/tools/cmd/godoc

echo "$0 done"
