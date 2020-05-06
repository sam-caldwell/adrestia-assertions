package assert
/*
	AdrestiaAssert (adrestia-assertion/pkg/assert.go)
	(c) 2020 Sam Caldwell.  See LICENSE.txt

	An argument spec is an object which represents a single object's grammar for use by the parser
    to represent the argument and to parse inputs into the final Argument object.
*/

import (
	"github.com/pkg/errors"
)

// Panic tests condition and panics if true.
func Panic(condition bool, response string){

	if condition{
		panic(response)
	}
}

// Error tests condition and returns error if true.
func Error(condition bool, response string) error {
	/*

	 */
	if condition{
		err:=errors.New(response)
		return err
	}
	return nil
}