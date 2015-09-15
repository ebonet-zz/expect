package goexpect

import "testing"

func Condition(t *testing.T, condition bool, description string) {
	if !condition {
		t.Error(description)
	}
}

func Equal(t *testing.T, expected interface{}, received interface{}, description string) {
	if expected != received {
		t.Errorf("[TEST ERROR] %s : Expected %v, Got %v",
			description,
			expected,
			received)
	}
}

func NoErr(t *testing.T, err error, description string) {
	if err != nil {
		t.Errorf("[TEST ERROR] %s : returned error: %s",
			description,
			err)
	}
}

func Err(t *testing.T, err error, description string) {
	if err == nil {
		t.Errorf("[TEST ERROR] %s : exepected to return error, but it didn't",
			description)
	}
}

func Between(t *testing.T, min float64, max float64, received float64, description string) {
	if min > received || max < received {
		t.Errorf("[TEST ERROR] %s : Expected %f < x < %f, Got %f",
			description,
			min,
			max,
			received)
	}
}

func Approximate(t *testing.T, expected float64, received float64, precision float64, description string) {
	Between(t, expected-precision, expected+precision, received, description)
}
