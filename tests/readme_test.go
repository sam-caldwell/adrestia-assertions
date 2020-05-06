/*
	AdrestiaAssert (license_text.go)
	(c) 2020 Sam Caldwell.  See LICENSE.txt

	Test the readme file
*/
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	Testing "testing"
)

func TestReadMeHash(t *Testing.T) {
	/*
		Hash the License file.  Ensure it matches the expected hash.
	*/
	const licenseFile = "../README.md"
	const expectedHash = "270e5ea63f22489c0bac9bc669d41a00438fd629a3a1579d3fdf8432c1817710"
	fileContent, err := ioutil.ReadFile(licenseFile)
	if err != nil {
		t.Errorf("File failed hash validation '%s'(%s): %v", licenseFile, expectedHash, err)
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(fileContent))
	if hash != expectedHash {
		t.Errorf("File hash mismatch: %s (%s)(%s)", licenseFile, expectedHash, hash)
	}
}

// ToDo: add more tests for individual documentation files.