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
	Assert.Panic(false, "No Panic")
	if Assert.Error(false, "No Error") == nil {
		if Assert.Error(true, "Expect error") != nil {
			Assert.Panic(true, "Now we panic")
		}
	}
}

func TestAssertPanicTrue(t *Testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic.")
		}
	}()
	Assert.Panic(true, "We should panic.")
}

func TestAssertPanicFalse(t *Testing.T) {
	Assert.Panic(false, "Unexpected panic.")
}

func TestAssertErrorTrue(t *Testing.T) {
	response := "Expected error"
	err := Assert.Error(true, response)
	if err == nil {
		t.Error("Expected error.  None returned.")
	}
	if fmt.Sprintf("%s", err) != response {
		t.Error("Expected error did not contain response.")
	}
}

func TestAssertErrorFalse(t *Testing.T) {
	err := Assert.Error(false, "Unexpected Error")
	if err != nil {
		t.Error("Expected error.  None returned.")
	}
}
