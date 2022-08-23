package validator

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

var (
	reUKPostCode = regexp.MustCompile(`^[a-zA-Z]{1,2}\d[a-zA-Z\d]?\s*\d[a-zA-Z]{2}$`)
	reZipCode    = regexp.MustCompile(`^(\d{5}(?:\-\d{4})?)$`)
)

const (
	validateEmpty       = "value cannot be empty"
	validateNotEmpty    = "value must be empty"
	validateLength      = "value must be between %d and %d characters"
	validateExactLength = "value should be exactly %d characters"
	validateMin         = "value %v is smaller than minimum %v"
	validateMax         = "value %v is larger than maximum %v"
	validateNumBetween  = "value %v must be between %v and %v"
	validatePositive    = "value %v should be greater than 0"
	validateRegex       = "value %s failed to meet requirements"
	validateBool        = "value %v does not evaluate to %v"
	validateDateEqual   = "the date/time provided %s, does not match the expected %s"
	validateDateAfter   = "the date provided %s, must be after %s"
	validateDateBefore  = "the date provided %s, must be before %s"
	validateUkPostCode  = "%s is not a valid UK PostCode"
	validateIsNumeric   = "string %s is not a number"
	validateEmail       = "invalid email"
)

// StrLength will ensure a string, val, has a length that is at least min and
// at most max.
func StrLength(val string, min, max int) ValidationFunc {
	return func() error {
		if len(val) >= min && len(val) <= max {
			return nil
		}
		return fmt.Errorf(validateLength, min, max)
	}
}

// StrLengthExact will ensure a string, val, is exactly length.
func StrLengthExact(val string, length int) ValidationFunc {
	return func() error {
		if len(val) == length {
			return nil
		}
		return fmt.Errorf(validateExactLength, length)
	}
}

// Number defines all number types.
type Number interface {
	constraints.Integer | constraints.Float
}

// MinNumber will ensure a Number, val, is at least min in value.
func MinNumber[T Number](val, min T) ValidationFunc {
	return func() error {
		if val >= min {
			return nil
		}
		return fmt.Errorf(validateMin, val, min)
	}
}

// MaxNumber will ensure an Int, val,  is at most Max in value.
func MaxNumber[T Number](val, max T) ValidationFunc {
	return func() error {
		if val <= max {
			return nil
		}
		return fmt.Errorf(validateMax, val, max)
	}
}

// BetweenNumber will ensure an int, val, is at least min and at most max.
func BetweenNumber[T Number](val, min, max T) ValidationFunc {
	return func() error {
		if val >= min && val <= max {
			return nil
		}
		return fmt.Errorf(validateNumBetween, val, min, max)
	}
}

// PositiveNumber will ensure an int, val, is > 0.
func PositiveNumber[T Number](val T) ValidationFunc {
	return func() error {
		if val > 0 {
			return nil
		}
		return fmt.Errorf(validatePositive, val)
	}
}

// MatchString will check that a string, val, matches the provided regular expression.
func MatchString(val string, r *regexp.Regexp) ValidationFunc {
	return func() error {
		if r.MatchString(val) {
			return nil
		}
		return fmt.Errorf(validateRegex, val)
	}
}

// MatchBytes will check that a byte array, val, matches the provided regular expression.
func MatchBytes(val []byte, r *regexp.Regexp) ValidationFunc {
	return func() error {
		if r.Match(val) {
			return nil
		}
		return fmt.Errorf(validateRegex, val)
	}
}

// Equal is a simple check to ensure that val matches exp.
func Equal[T comparable](val, exp T) ValidationFunc {
	return func() error {
		if val == exp {
			return nil
		}
		return fmt.Errorf(validateBool, val, exp)
	}
}

// DateEqual will ensure that a date/time, val, matches exactly exp.
func DateEqual(val, exp time.Time) ValidationFunc {
	return func() error {
		if val.Equal(exp) {
			return nil
		}
		return fmt.Errorf(validateDateEqual, val, exp)
	}
}

// DateAfter will ensure that a date/time, val, occurs after exp.
func DateAfter(val, exp time.Time) ValidationFunc {
	return func() error {
		if val.After(exp) {
			return nil
		}
		return fmt.Errorf(validateDateAfter, val, exp)
	}
}

// DateBefore will ensure that a date/time, val, occurs before exp.
func DateBefore(val, exp time.Time) ValidationFunc {
	return func() error {
		if val.Before(exp) {
			return nil
		}
		return fmt.Errorf(validateDateBefore, val, exp)
	}
}

// NotEmpty will ensure that a value, val, is not empty.
// rules are:
// int: > 0
// string: != "" or whitespace
// slice: not nil and len > 0
// map: not nil and len > 0
func NotEmpty(v interface{}) ValidationFunc {
	return func() error {
		if v == nil {
			return fmt.Errorf(validateEmpty)
		}
		val := reflect.ValueOf(v)
		valid := false
		// nolint:exhaustive // not supporting everything
		switch val.Kind() {
		case reflect.Map, reflect.Slice:
			valid = val.Len() > 0 && !val.IsNil()
		default:
			valid = !val.IsZero()
		}
		if !valid {
			return fmt.Errorf(validateEmpty)
		}
		return nil
	}
}

// Empty will ensure that a value, val, is empty.
// rules are:
// int: == 0
// string: == "" or whitespace
// slice: is nil or len == 0
// map: is nil and len == 0
func Empty(v interface{}) ValidationFunc {
	return func() error {
		err := NotEmpty(v)()
		if err == nil {
			return fmt.Errorf(validateNotEmpty)
		}
		return nil
	}
}

// IsNumeric will pass if a string, val, is an Int.
func IsNumeric(val string) ValidationFunc {
	return func() error {
		_, err := strconv.Atoi(val)
		if err == nil {
			return nil
		}
		return fmt.Errorf(validateIsNumeric, val)
	}
}

// UKPostCode will validate that a string, val, is a valid UK PostCode.
// It does not check the postcode exists, just that it matches an agreed pattern.
func UKPostCode(val string) ValidationFunc {
	return func() error {
		if reUKPostCode.MatchString(val) {
			return nil
		}
		return fmt.Errorf(validateUkPostCode, val)
	}
}

// USZipCode will validate that a string, val, matches a US USZipCode pattern.
// It does not check the zipcode exists, just that it matches an agreed pattern.
func USZipCode(val string) ValidationFunc {
	return func() error {
		if reZipCode.MatchString(val) {
			return nil
		}
		return fmt.Errorf("%s is not a valid UK PostCode", val)
	}
}

// HasPrefix ensures string, val, has a prefix matching prefix.
func HasPrefix(val, prefix string) ValidationFunc {
	return func() error {
		if strings.HasPrefix(val, prefix) {
			return nil
		}
		return fmt.Errorf("value provided does not have a valid prefix")
	}
}

// NoPrefix ensures a string, val, does not have the supplied prefix.
func NoPrefix(val, prefix string) ValidationFunc {
	return func() error {
		if strings.HasPrefix(val, prefix) {
			return errors.New("value provided does not have a valid prefix")
		}
		return nil
	}
}

// IsHex will check that a string, val, is valid Hexadecimal.
func IsHex(val string) ValidationFunc {
	return func() error {
		if _, err := hex.DecodeString(val); err != nil {
			return errors.New("value supplied is not valid hex")
		}
		return nil
	}
}

// Email will check that a string is a valid email address.
func Email(val string) ValidationFunc {
	return func() error {
		if _, err := mail.ParseAddress(val); err != nil {
			return errors.New(validateEmail)
		}
		return nil
	}
}

// AnyString will check if the provided string is in a set of allowed values.
func AnyString(val string, vv ...string) ValidationFunc {
	return func() error {
		for _, v := range vv {
			if val == v {
				return nil
			}
		}

		return errors.New("value not found in allowed values")
	}
}
