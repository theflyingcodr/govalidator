package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	validateEmpty  = "value cannot be empty"
	validateLength = "value %s must be between %d and %d"
)

// Length will ensure a string has a length that is at least min and
// at most max.
func Length(val string, min, max int) ValidationFunc {
	return func() error {
		if len(val) >= min && len(val) <= max {
			return nil
		}
		return fmt.Errorf(validateLength, val, min, max)
	}
}

// MinInt will ensure an Int is at least min in value.
func MinInt(val, min int) ValidationFunc {
	return func() error {
		if val >= min {
			return nil
		}
		return fmt.Errorf("value %d is smaller than minimum %d", val, min)
	}
}

// MaxInt will ensure an Int, val,  is at most Max in value.
func MaxInt(val, max int) ValidationFunc {
	return func() error {
		if val <= max {
			return nil
		}
		return fmt.Errorf("value %d is larger than maximum %d", val, max)
	}
}

// BetweenInt will ensure an int, val,  is at least min and at most max.
func BetweenInt(val, min, max int) ValidationFunc {
	return func() error {
		if val >= min && val <= max {
			return nil
		}
		return fmt.Errorf("value %d must be between %d and %d", val, min, max)
	}
}

// MinInt64 will ensure an Int64, val, is at least min in value.
func MinInt64(val, min int) ValidationFunc {
	return func() error {
		if val >= min {
			return nil
		}
		return fmt.Errorf("value %d is smaller than minimum %d", val, min)
	}
}

// MaxInt64 will ensure an Int64, val, is at most Max in value.
func MaxInt64(val, max int) ValidationFunc {
	return func() error {
		if val <= max {
			return nil
		}
		return fmt.Errorf("value %d is larger than maximum %d", val, max)
	}
}

// BetweenInt64 will ensure an int64, val, is at least min and at most max.
func BetweenInt64(val, min, max int64) ValidationFunc {
	return func() error {
		if val >= min && val <= max {
			return nil
		}
		return fmt.Errorf("value %d must be between %d and %d", val, min, max)
	}
}

func PositiveInt(val int) ValidationFunc {
	return func() error {
		if val > 0 {
			return nil
		}
		return fmt.Errorf("value %d should be greater than 0", val)
	}
}

func PositiveInt64(val int64) ValidationFunc {
	return func() error {
		if val > 0 {
			return nil
		}
		return fmt.Errorf("value %d should be greater than 0", val)
	}
}

// Match will check that a string, val, matches the provided regular expression.
func Match(val string, r *regexp.Regexp) ValidationFunc {
	return func() error {
		if r.MatchString(val) {
			return nil
		}
		return fmt.Errorf("value %s failed to meet requirements", val)
	}
}

// MatchBytes will check that a byte array, val, matches the provided regular expression.
func MatchBytes(val []byte, r *regexp.Regexp) ValidationFunc {
	return func() error {
		if r.Match(val) {
			return nil
		}
		return fmt.Errorf("value %s failed to meet requirements", val)
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
		switch val.Kind() {
		case reflect.Array, reflect.Map, reflect.Slice:
			valid = val.Len() > 0 && !val.IsNil()
		case reflect.String:
			valid = len(strings.TrimSpace(val.String())) > 0
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valid = val.Int() > 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			valid = val.Uint() > 0
		case reflect.Float32, reflect.Float64:
			valid = val.Float() > 0
		case reflect.Interface, reflect.Ptr:
			valid = !val.IsNil()
		default:
			panic(fmt.Errorf("unsupported type %T", v))
		}
		if !valid {
			return fmt.Errorf(validateEmpty)
		}
		return nil
	}
}
