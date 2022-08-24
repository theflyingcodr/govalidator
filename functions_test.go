package validator

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestLength(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		s      string
		minLen int
		maxLen int
		expErr error
	}{
		"string within length bounds": {
			s:      "hi there",
			minLen: 1,
			maxLen: 100,
		},
		"string at minimum length": {
			s:      "hi there's",
			minLen: 10,
			maxLen: 100,
		},
		"string at max length": {
			s:      "hi there",
			minLen: 1,
			maxLen: 8,
		},
		"string too small": {
			s:      "hi there",
			minLen: 50,
			maxLen: 80,
			expErr: fmt.Errorf(validateLength, 50, 80),
		},
		"string too large": {
			s:      "hi there",
			minLen: 1,
			maxLen: 4,
			expErr: fmt.Errorf(validateLength, 1, 4),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, StrLength(test.s, test.minLen, test.maxLen)())
		})
	}
}

func TestMinInt(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int
		min    int
		expErr error
	}{
		"int larger than min should pass": {
			i:   10,
			min: 5,
		},
		"int equal to min should pass": {
			i:   5,
			min: 5,
		},
		"int smaller than min should fail": {
			i:      5,
			min:    50,
			expErr: fmt.Errorf(validateMin, 5, 50),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, MinNumber(test.i, test.min)())
		})
	}
}

func TestMaxInt(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int
		max    int
		expErr error
	}{
		"int smaller than max should pass": {
			i:   100,
			max: 101,
		},
		"int equal to max should pass": {
			i:   5,
			max: 5,
		},
		"int larger than max should fail": {
			i:      51,
			max:    50,
			expErr: fmt.Errorf(validateMax, 51, 50),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, MaxNumber(test.i, test.max)())
		})
	}
}

func TestBetweenInt(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int
		min    int
		max    int
		expErr error
	}{
		"int between min and max should pass": {
			i:   100,
			min: 50,
			max: 101,
		},
		"int equal to max should pass": {
			i:   100,
			min: 50,
			max: 100,
		},
		"int equal to min should pass": {
			i:   50,
			min: 50,
			max: 100,
		},
		"int larger than max should fail": {
			i:      51,
			max:    50,
			expErr: fmt.Errorf(validateNumBetween, 51, 0, 50),
		},
		"int smaller than min should fail": {
			i:      5,
			min:    6,
			max:    50,
			expErr: fmt.Errorf(validateNumBetween, 5, 6, 50),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, BetweenNumber(test.i, test.min, test.max)())
		})
	}
}

func TestMinInt64(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int64
		min    int64
		expErr error
	}{
		"int larger than min should pass": {
			i:   10,
			min: 5,
		},
		"int equal to min should pass": {
			i:   5,
			min: 5,
		},
		"int smaller than min should fail": {
			i:      5,
			min:    50,
			expErr: fmt.Errorf(validateMin, 5, 50),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, MinNumber(test.i, test.min)())
		})
	}
}

func TestMaxInt64(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int64
		max    int64
		expErr error
	}{
		"int smaller than max should pass": {
			i:   100,
			max: 101,
		},
		"int equal to max should pass": {
			i:   5,
			max: 5,
		},
		"int larger than max should fail": {
			i:      51,
			max:    50,
			expErr: fmt.Errorf(validateMax, 51, 50),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, MaxNumber(test.i, test.max)())
		})
	}
}

func TestBetweenInt64(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int64
		min    int64
		max    int64
		expErr error
	}{
		"int between min and max should pass": {
			i:   100,
			min: 50,
			max: 101,
		},
		"int equal to max should pass": {
			i:   100,
			min: 50,
			max: 100,
		},
		"int equal to min should pass": {
			i:   50,
			min: 50,
			max: 100,
		},
		"int larger than max should fail": {
			i:      51,
			max:    50,
			expErr: fmt.Errorf(validateNumBetween, 51, 0, 50),
		},
		"int smaller than min should fail": {
			i:      5,
			min:    6,
			max:    50,
			expErr: fmt.Errorf(validateNumBetween, 5, 6, 50),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, BetweenNumber(test.i, test.min, test.max)())
		})
	}
}

func TestPositiveInt(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int
		expErr error
	}{
		"int greater than 0 should pass": {
			i: 100,
		},
		"int max should pass": {
			i: 2147483647,
		},
		"int smaller than 0 should fail": {
			i:      -1,
			expErr: fmt.Errorf(validatePositive, -1),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, PositiveNumber(test.i)())
		})
	}
}

func TestPositiveInt64(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		i      int64
		expErr error
	}{
		"int64 greater than 0 should pass": {
			i: 100,
		},
		"int64 max should pass": {
			i: 9223372036854775807,
		},
		"int64 smaller than 0 should fail": {
			i:      -1,
			expErr: fmt.Errorf(validatePositive, -1),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, PositiveNumber(test.i)())
		})
	}
}

func TestMatchString(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		s      string
		r      *regexp.Regexp
		expErr error
	}{
		"string that matches should pass": {
			s: "hi there",
			r: regexp.MustCompile("[a-z ]*"),
		},
		"string that matches list should pass": {
			s: "pass",
			r: regexp.MustCompile(`(pass|fail)`),
		},
		"string that doesn't match should fail": {
			s:      "oops",
			r:      regexp.MustCompile(`(pass|fail)`),
			expErr: fmt.Errorf(validateRegex, "oops"),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, MatchString(test.s, test.r)())
		})
	}
}

func TestMatchBytes(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		s      []byte
		r      *regexp.Regexp
		expErr error
	}{
		"string that matches should pass": {
			s: []byte("hi there"),
			r: regexp.MustCompile("[a-z ]*"),
		},
		"string that matches list should pass": {
			s: []byte("pass"),
			r: regexp.MustCompile(`(pass|fail)`),
		},
		"string that doesn't match should fail": {
			s:      []byte("oops"),
			r:      regexp.MustCompile(`(pass|fail)`),
			expErr: fmt.Errorf(validateRegex, "oops"),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, MatchBytes(test.s, test.r)())
		})
	}
}
func TestEqualBool(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    bool
		exp    bool
		expErr error
	}{
		"val matching exp should pass": {
			val: true,
			exp: true,
		}, "val matching false exp should pass": {
			val: false,
			exp: false,
		}, "val not matching exp should fail": {
			val:    true,
			exp:    false,
			expErr: fmt.Errorf(validateBool, true, false),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, Equal(test.val, test.exp)())
		})
	}
}

func TestEqualString(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    string
		exp    string
		expErr error
	}{
		"val matching exp should pass": {
			val: "hi there, this is a test!",
			exp: "hi there, this is a test!",
		}, "val not matching exp should fail": {
			val:    "hi there, this is a test",
			exp:    "hi there, this is a test! but i'm different",
			expErr: fmt.Errorf(validateBool, "hi there, this is a test", "hi there, this is a test! but i'm different"),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, Equal(test.val, test.exp)())
		})
	}
}

func TestEqualInt(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    int
		exp    int
		expErr error
	}{
		"val matching exp should pass": {
			val: 1234,
			exp: 1234,
		}, "val not matching exp should fail": {
			val:    1234,
			exp:    433321,
			expErr: fmt.Errorf(validateBool, 1234, 433321),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, Equal(test.val, test.exp)())
		})
	}
}

func TestEqualDate(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    time.Time
		exp    time.Time
		expErr error
	}{
		"val matching exp should pass": {
			val: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
		}, "val not matching exp should fail": {
			val: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 2, time.UTC),
			expErr: fmt.Errorf(validateBool,
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
				time.Date(2021, 1, 1, 1, 1, 1, 2, time.UTC)),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, Equal(test.val, test.exp)())
		})
	}
}

func TestDateEqual(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    time.Time
		exp    time.Time
		expErr error
	}{
		"date matching should pass": {
			val: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		"date not matching should fail": {
			val: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 2, time.UTC),
			expErr: fmt.Errorf(validateDateEqual,
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
				time.Date(2021, 1, 1, 1, 1, 1, 2, time.UTC)),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, DateEqual(test.val, test.exp)())
		})
	}
}

func TestDateBefore(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    time.Time
		exp    time.Time
		expErr error
	}{
		"date before should pass": {
			val: time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		"date matching exp should fail": {
			val: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			expErr: fmt.Errorf(validateDateBefore,
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)),
		},
		"date after exp should fail": {
			val: time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			expErr: fmt.Errorf(validateDateBefore,
				time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC),
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, DateBefore(test.val, test.exp)())
		})
	}
}

func TestDateAfter(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    time.Time
		exp    time.Time
		expErr error
	}{
		"date after should pass": {
			val: time.Date(2021, 1, 1, 1, 1, 1, 2, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		"date matching exp should fail": {
			val: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			expErr: fmt.Errorf(validateDateAfter,
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)),
		},
		"date before exp should fail": {
			val: time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC),
			exp: time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
			expErr: fmt.Errorf(validateDateAfter,
				time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC),
				time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, DateAfter(test.val, test.exp)())
		})
	}
}

func TestIsNumeric(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    string
		expErr error
	}{
		"valid number should pass": {
			val: "12345",
		},
		"valid negative number should pass": {
			val: "-12345",
		},
		"invalid number should fail": {
			val:    "12345a",
			expErr: fmt.Errorf(validateIsNumeric, "12345a"),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, IsNumeric(test.val)())
		})
	}
}

func TestUKPostCode(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    []string
		expErr error
	}{
		"Valid postcodes should pass": {
			val: []string{"bt13 4GH", "NW1A 1AA", "A9A 9AA", "A9 9AA", "A99 9AA"},
		},
		"Invalid postcodes should fail": {
			val:    []string{"GGG 7GH", "NW1A 1A", "N1 GF", "N11 DDD"},
			expErr: fmt.Errorf(validateUkPostCode, "GGG 7GH"),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			for _, p := range test.val {
				err := UKPostCode(p)()
				if test.expErr == nil {
					is.NoErr(err)
					continue
				}
				is.Equal(err != nil, true)
			}
		})
	}
}

func TestZipCode(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    []string
		expErr error
	}{
		"Valid zipcodes should pass": {
			val: []string{"57501", "17101", "12201-7050", "99750-0077"},
		},
		"Invalid zipcodes should fail": {
			val:    []string{"GGG 7GH", "99750-00", "99750-0", "99750-", "1111"},
			expErr: fmt.Errorf(validateUkPostCode, "GGG 7GH"),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			for _, p := range test.val {
				err := USZipCode(p)()
				if test.expErr == nil {
					is.NoErr(err)
					continue
				}
				is.Equal(err != nil, true)
			}
		})
	}
}

func TestEmail(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    string
		expErr error
	}{
		"email without a domain should fail": {
			val:    "test@",
			expErr: fmt.Errorf(validateEmail),
		},
		"email without a prefix": {
			val:    "@test.com",
			expErr: fmt.Errorf(validateEmail),
		},
		"emails are not required to have a tld so will pass": {
			val: "test@mail",
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, Email(test.val)())
		})
	}
}

func TestNotEmpty(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    interface{}
		expErr error
	}{
		"nil ptr": {
			val:    nil,
			expErr: errors.New(validateEmpty),
		},
		"non-empty string": {
			val: "hello",
		},
		"empty string": {
			val:    "",
			expErr: errors.New(validateEmpty),
		},
		"non-empty time": {
			val: time.Now(),
		},
		"empty time": {
			val:    time.Time{},
			expErr: errors.New(validateEmpty),
		},
		"non-empty int": {
			val: 235,
		},
		"empty int": {
			val:    0,
			expErr: errors.New(validateEmpty),
		},
		"non-empty int8": {
			val: int8(5),
		},
		"empty int8": {
			val:    int8(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty int16": {
			val: int16(5),
		},
		"empty int16": {
			val:    int16(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty int32": {
			val: int32(5),
		},
		"empty int32": {
			val:    int32(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty int64": {
			val: int64(5),
		},
		"empty int64": {
			val:    int64(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty uint": {
			val: 235,
		},
		"empty uint": {
			val:    0,
			expErr: errors.New(validateEmpty),
		},
		"non-empty uint8": {
			val: uint8(5),
		},
		"empty uint8": {
			val:    uint8(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty uint16": {
			val: uint16(5),
		},
		"empty uint16": {
			val:    uint16(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty uint32": {
			val: uint32(5),
		},
		"empty uint32": {
			val:    uint32(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty uint64": {
			val: uint64(5),
		},
		"empty uint64": {
			val:    uint64(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty float32": {
			val: float32(5),
		},
		"empty float32": {
			val:    float32(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty float64": {
			val: float64(5),
		},
		"empty float64": {
			val:    float64(0),
			expErr: errors.New(validateEmpty),
		},
		"non-empty array": {
			val: [2]string{"hello", "there"},
		},
		"empty array": {
			val:    [2]string{"", ""},
			expErr: errors.New(validateEmpty),
		},
		"non-empty slice": {
			val: []string{"hello", "there"},
		},
		"empty slice": {
			val:    []string{},
			expErr: errors.New(validateEmpty),
		},
		"non-empty map": {
			val: map[string]string{"hello": "there"},
		},
		"empty map": {
			val:    map[string]string{},
			expErr: errors.New(validateEmpty),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, NotEmpty(test.val)())
		})
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    interface{}
		expErr error
	}{
		"nil ptr": {
			val: nil,
		},
		"non-empty string": {
			val:    "hello",
			expErr: errors.New(validateNotEmpty),
		},
		"empty string": {
			val: "",
		},
		"non-empty time": {
			val:    time.Now(),
			expErr: errors.New(validateNotEmpty),
		},
		"empty time": {
			val: time.Time{},
		},
		"non-empty int": {
			val:    235,
			expErr: errors.New(validateNotEmpty),
		},
		"empty int": {
			val: 0,
		},
		"non-empty int8": {
			val:    int8(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty int8": {
			val: int8(0),
		},
		"non-empty int16": {
			val:    int16(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty int16": {
			val: int16(0),
		},
		"non-empty int32": {
			val:    int32(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty int32": {
			val: int32(0),
		},
		"non-empty int64": {
			val:    int64(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty int64": {
			val: int64(0),
		},
		"non-empty uint": {
			val:    235,
			expErr: errors.New(validateNotEmpty),
		},
		"empty uint": {
			val: 0,
		},
		"non-empty uint8": {
			val:    uint8(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty uint8": {
			val: uint8(0),
		},
		"non-empty uint16": {
			val:    uint16(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty uint16": {
			val: uint16(0),
		},
		"non-empty uint32": {
			val:    uint32(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty uint32": {
			val: uint32(0),
		},
		"non-empty uint64": {
			val:    uint64(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty uint64": {
			val: uint64(0),
		},
		"non-empty float32": {
			val:    float32(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty float32": {
			val: float32(0),
		},
		"non-empty float64": {
			val:    float64(5),
			expErr: errors.New(validateNotEmpty),
		},
		"empty float64": {
			val: float64(0),
		},
		"non-empty array": {
			val:    [2]string{"hello", "there"},
			expErr: errors.New(validateNotEmpty),
		},
		"empty array": {
			val: [2]string{"", ""},
		},
		"non-empty slice": {
			val:    []string{"hello", "there"},
			expErr: errors.New(validateNotEmpty),
		},
		"empty slice": {
			val: []string{},
		},
		"non-empty map": {
			val:    map[string]string{"hello": "there"},
			expErr: errors.New(validateNotEmpty),
		},
		"empty map": {
			val: map[string]string{},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, Empty(test.val)())
		})
	}
}

func TestAnyString(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tt := map[string]struct {
		val    string
		list   []string
		expErr error
	}{
		"matching item": {
			val:  "hello",
			list: []string{"wow", "hello", "ohwow"},
		},
		"missing item": {
			val:    "hello",
			list:   []string{"wow", "goodbye", "ohwow"},
			expErr: errors.New("value not found in allowed values"),
		},
		"empty string found": {
			val:  "",
			list: []string{"wow", "", "ohwow"},
		},
		"empty string not found": {
			val:    "",
			list:   []string{"wow", "ohwow"},
			expErr: errors.New("value not found in allowed values"),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			is.Equal(test.expErr, AnyString(test.val, test.list...)())
		})
	}
}
