package validator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func Test_NewSingleError(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct {
		fieldName string
		errors    []string
	}{
		"empty string should return error with empty string": {
			fieldName: "test",
			errors:    []string{},
		}, "single string should return error with single string": {
			fieldName: "test",
			errors:    []string{"i failed"},
		}, "multiple string should return error with multiple string": {
			fieldName: "test",
			errors:    []string{"i failed", "i failed because of another thing"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := NewSingleError(test.fieldName, test.errors)
			var val ErrValidation
			ok := errors.As(err, &val)
			is = is.NewRelaxed(t)
			is.True(ok)
			is.Equal(val[test.fieldName], test.errors)
		})
	}
}

func Test_NewFromError(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct {
		fieldName string
		error     error
		exp       ErrValidation
	}{
		"nil error string should return error with empty string": {
			fieldName: "test",
			error:     nil,
			exp:       nil,
		}, "error should be returned in validator": {
			fieldName: "test",
			error:     errors.New("I failed"),
			exp: map[string][]string{
				"test": {"I failed"},
			},
		}, "wrapped error should be returned in validator": {
			fieldName: "test",
			error:     fmt.Errorf("i wrap %w", errors.New("I failed")),
			exp: map[string][]string{
				"test": {"i wrap I failed"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := NewFromError(test.fieldName, test.error)
			//val, ok := err.(ErrValidation)
			is = is.NewRelaxed(t)
			is.Equal(test.exp, err)
		})
	}
}
