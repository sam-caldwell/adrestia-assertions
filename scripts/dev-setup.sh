#!/bin/bash

echo "Install GoLint"
go get -u golang.org/x/lint/golint

echo "Install GoDoc"
go get golang.org/x/tools/cmd/godoc

echo "Install yamllint"
brew install yamllint

echo "$0 done"
