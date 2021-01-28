package validator

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/matryer/is"
)

func TestLength(t *testing.T) {
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
			expErr: fmt.Errorf(validateLength, "hi there", 50, 80),
		},
		"string too large": {
			s:      "hi there",
			minLen: 1,
			maxLen: 4,
			expErr: fmt.Errorf(validateLength, "hi there", 1, 4),
		},
	}
	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, Length(test.s, test.minLen, test.maxLen)())
		})
	}
}

func TestMinInt(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, MinInt(test.i, test.min)())
		})
	}
}

func TestMaxInt(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, MaxInt(test.i, test.max)())
		})
	}
}

func TestBetweenInt(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, BetweenInt(test.i, test.min, test.max)())
		})
	}
}

func TestMinInt64(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, MinInt64(test.i, test.min)())
		})
	}
}

func TestMaxInt64(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, MaxInt64(test.i, test.max)())
		})
	}
}

func TestBetweenInt64(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, BetweenInt64(test.i, test.min, test.max)())
		})
	}
}

func TestPositiveInt(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, PositiveInt(test.i)())
		})
	}
}

func TestPositiveInt64(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, PositiveInt64(test.i)())
		})
	}
}

func TestMatchString(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, MatchString(test.s, test.r)())
		})
	}
}

func TestMatchBytes(t *testing.T) {
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
			is := is.NewRelaxed(t)
			is.Equal(test.expErr, MatchBytes(test.s, test.r)())
		})
	}
}
