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
	const expectedHash = "0eef3fab105c796e1ea0a291b31fed9755384569dcbbec69fa06377ffbc53063"
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