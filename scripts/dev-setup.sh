#!/bin/bash

export GO111MODULE="on"

echo "Install git-hooks"
go get -u github.com/git-hooks/git-hooks

echo "Install the git hooks to the local environment."
git hooks install

echo "List current hooks"
git hooks list
echo "$0 done"
