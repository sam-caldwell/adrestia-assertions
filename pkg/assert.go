/*
	AdrestiaAssert (adrestia-assertion/pkg/assert.go)
	(c) 2020 Sam Caldwell.  See LICENSE.txt

	An argument spec is an object which represents a single object's grammar for use by the parser
    to represent the argument and to parse inputs into the final Argument object.
*/
package AdrestiaAssert

import (
	"github.com/pkg/errors"
)

func Panic(condition bool, response string){
	if condition{
		panic(response)
	}
}
func Error(condition bool, response string) (error) {
	if condition{
		return errors.New(response)
	}
	return nil
}