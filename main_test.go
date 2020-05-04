/*
	AdrestiaAssert (adrestia-assertion/main_test.go)
	(c) 2020 Sam Caldwell.  See LICENSE.txt

	Test the assert functions
*/
package main

import (
	"crypto/sha256"
	"fmt"
	AdrestiaAssert "github.com/sam-caldwell/adrestia-assertions/pkg"
	"io/ioutil"
	Testing "testing"
)

func main() {
	AdrestiaAssert.Panic(false, "No Panic")
	if AdrestiaAssert.Error(false, "No Error") == nil {
		if AdrestiaAssert.Error(true, "Expect error") != nil {
			AdrestiaAssert.Panic(true, "Now we panic")
		}
	}
}

func TestAssertPanicTrue(t *Testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic.")
		}
	}()
	AdrestiaAssert.Panic(true, "We should panic.")
}

func TestAssertPanicFalse(t *Testing.T) {
	AdrestiaAssert.Panic(false, "Unexpected panic.")
}

func TestAssertErrorTrue(t *Testing.T) {
	response := "Expected error"
	err := AdrestiaAssert.Error(true, response)
	if err == nil {
		t.Error("Expected error.  None returned.")
	}
	if fmt.Sprintf("%s", err) != response {
		t.Error("Expected error did not contain response.")
	}
}

func TestAssertErrorFalse(t *Testing.T) {
	err := AdrestiaAssert.Error(false, "Unexpected Error")
	if err != nil {
		t.Error("Expected error.  None returned.")
	}
}

func TestLicenseHash(t *Testing.T) {
	/*
		Hash the License file.  Ensure it matches the expected hash.
	*/
	const licenseFile = "LICENSE.txt"
	const expectedHash = "72c552ad631e5d855b6e57ad38afce1dbec6807aec98232d599fb821d5285fcf"
	fileContent, err := ioutil.ReadFile(licenseFile)
	if err != nil {
		t.Errorf("File failed hash validation '%s'(%s): %v", licenseFile, expectedHash, err)
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(fileContent))
	if hash != expectedHash {
		t.Errorf("File hash mismatch: %s (%s)(%s)",licenseFile,expectedHash,hash)
	}
}

func TestReadMeHash(t *Testing.T) {
	/*
		Hash the License file.  Ensure it matches the expected hash.
	*/
	const licenseFile = "README.md"
	const expectedHash = "270e5ea63f22489c0bac9bc669d41a00438fd629a3a1579d3fdf8432c1817710"
	fileContent, err := ioutil.ReadFile(licenseFile)
	if err != nil {
		t.Errorf("File failed hash validation '%s'(%s): %v", licenseFile, expectedHash, err)
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(fileContent))
	if hash != expectedHash {
		t.Errorf("File hash mismatch: %s (%s)(%s)",licenseFile,expectedHash,hash)
	}
}
