# Makefile
# (c) 2020 Sam Caldwell.  See LICENSE.txt
#
setup:
	./scripts/dev-setup.sh

lint:
	git hooks run pre-commit

go_test:
	git hooks run pre-push

test: go_test

