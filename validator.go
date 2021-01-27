package validator

import (
	"fmt"
	"sort"
	"strings"
)

type Validator interface {
	Validate() error
}

type ValidationFunc func() error

func (v ValidationFunc) String() string {
	return v().Error()
}

type ErrValidation map[string][]string

func New() ErrValidation {
	return map[string][]string{}
}

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
