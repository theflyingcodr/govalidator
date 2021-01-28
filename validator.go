package validator

import (
	"fmt"
	"sort"
	"strings"
)

type Validator interface {
	Validate() error
}

// ValidationFunc defines a simple function that can be wrapped
// and supplied with arguments.
type ValidationFunc func() error

// String satisfies the String interface and returns the underlying error
// string that is returned by evaluating the function.
func (v ValidationFunc) String() string {
	return v().Error()
}

// ErrValidation contains a list of field names and a list of errors
// found against each. This can then be converted for output to a user.
type ErrValidation map[string][]string

// New will create and return a new ErrValidation which can have Validate functions chained.
func New() ErrValidation {
	return map[string][]string{}
}

// Validate will log any errors found when evaluating the list of validation functions
// supplied to it
func (e ErrValidation) Validate(field string, fns ...ValidationFunc) ErrValidation {
	out := make([]string, len(fns), len(fns))
	for _, fn := range fns {
		if err := fn(); err != nil {
			out = append(out, err.Error())
		}
	}
	if len(out) > 0 {
		e[field] = out
	}
	return e
}

// IsValid will return true if no errors are found, ie all validators return valid
// or false if an error has been found.
func (e ErrValidation) IsValid() bool {
	return len(e) == 0
}

func (e ErrValidation) String() string {
	if e.IsValid() {
		return "no validation errors"
	}
	errs := make([]string, 0)
	for k, vv := range e {
		errs = append(errs, fmt.Sprintf("[%s: %s]", k, strings.Join(vv, ", ")))
	}
	sort.Strings(errs)
	return strings.Join(errs, ", ")
}

func (e ErrValidation) Error() string {
	return e.String()
}
