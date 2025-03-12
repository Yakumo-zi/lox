package errors

import "fmt"

var Errors []error = make([]error, 10)

func Report(line int, where, msg string) {
	Errors = append(Errors, fmt.Errorf("[line %d] Error %s:%s", line, where, msg))
}
