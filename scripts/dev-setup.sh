#!/bin/bash -e
export GO111MODULE="on"
export GOPATH="$GOPATH:$(pwd)/"

echo "GOPATH: $GOPATH"

echo "Install git-hooks"

#go get -u github.com/git-hooks/git-hooks

echo "Install the git hooks to the local environment."
(
    mkdir -p $GOPATH/src/github.com/git-hooks
    cd $GOPATH/src/github.com/git-hooks
    git clone git@github.com:git-hooks/git-hooks.git
    cd git-hooks/
    # install godep and restore deps
    make get
    # install binary
    go install
)
git hooks install

echo "List current hooks"
git hooks list

echo "Install GoDoc"
go get golang.org/x/tools/cmd/godoc

echo "Install yamllint"
brew install yamllint

echo "Install golint"
go get -u golang.org/x/lint/golint

echo "$0 done"