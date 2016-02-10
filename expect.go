package expect

import (
	"testing"
	"runtime"
	"strings"
	"fmt"
)

func Condition(t *testing.T, condition bool, description string) {
	if !condition {
		fail(t, description)
	}
}

func Equal(t *testing.T, expected interface{}, received interface{}, description string) {
	if expected != received {
		fail(t, "%s : Expected %v, Got %v",
			description,
			expected,
			received)
	}
}

func NoErr(t *testing.T, err error, description string) {
	if err != nil {
		fail(t, "%s : returned error: %s",
			description,
			err)
	}
}

func Err(t *testing.T, err error, description string) {
	if err == nil {
		fail(t, " %s : expected to return error, but it didn't",
			description)
	}
}

func Between(t *testing.T, min float64, max float64, received float64, description string) {
	if min > received || max < received {

		fail(t, "%s : Expected %f < x < %f, Got %f",
			description,
			min,
			max,
			received)
	}
}

func Approximate(t *testing.T, expected float64, received float64, precision float64, description string) {
	Between(t, expected-precision, expected+precision, received, description)
}

func In(t *testing.T, theMap map[interface{}]interface{}, key interface{}, description string) {
	if _, ok := theMap[key]; !ok {
		fail(t, "%s : Map does not contain key %v", description, key)
	}
}

func fail(t *testing.T, message string, params ...interface{}) {
	file, line := getLocation()
	t.Errorf("[TEST ERROR {%s:d}] "+message,
		append([]interface{}{file, line}, params...)...)
}


func getLocation() (file string, line int) {
	i := 1
	for ; i < 9; i++ {
		_, file, _, ok := runtime.Caller(i)
		if !ok || strings.Contains(file, "testing/testing.go") {
			break
		}
	}
	_, file, line, ok := runtime.Caller(i - 1)
	if ok && i > 1 && i < 9 {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	return file, line
}