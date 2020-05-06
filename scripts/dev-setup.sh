#!/bin/bash -e

echo "GOPATH: $GOPATH"

echo "Install yamllint"
brew install yamllint

echo "Install golint"
go get -u golang.org/x/lint/golint

echo "Install git-hooks"
export GO111MODULE="on"
go get -u github.com/git-hooks/git-hooks

echo "Install the git hooks to the local environment."
git hooks install

echo "List current hooks"
git hooks list

echo "Install GoDoc"
go get golang.org/x/tools/cmd/godoc

echo "$0 done"