/*
	AdrestiaAssert (adrestia-assertion/main_test.go)
	(c) 2020 Sam Caldwell.  See LICENSE.txt

	Test the assert functions
*/
package main

import (
	"fmt"
	"github.com/sam-caldwell/adrestia-assertions/src"
	Testing "testing"
)

func main() {
	assert.Panic(false, "No Panic")
	if assert.Error(false, "No Error") == nil {
		if assert.Error(true, "Expect error") != nil {
			assert.Panic(true, "Now we panic")
		}
	}
}

func TestAssertPanicTrue(t *Testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic.")
		}
	}()
	assert.Panic(true, "We should panic.")
}

func TestAssertPanicFalse(t *Testing.T) {
	assert.Panic(false, "Unexpected panic.")
}

func TestAssertErrorTrue(t *Testing.T) {
	response := "Expected error"
	err := assert.Error(true, response)
	if err == nil {
		t.Error("Expected error.  None returned.")
	}
	if fmt.Sprintf("%s", err) != response {
		t.Error("Expected error did not contain response.")
	}
}

func TestAssertErrorFalse(t *Testing.T) {
	err := assert.Error(false, "Unexpected Error")
	if err != nil {
		t.Error("Expected error.  None returned.")
	}
}
