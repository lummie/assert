package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"unicode/utf8"
)

// Equal tries to establish if the two values are compareEquality via reflection and if that fails then via conversion to string
func Equal(t *testing.T, expected, actual interface{}, message ...string) {
	if !compareEquality(expected, actual) {
		t.Errorf("%v Expected \n[%#v]\nto be\n[%#v]\n%v ", message, expected, actual, callerInfo(2))
	}
}

// NotEqual utilises the same method as Equal but returns the complement
func NotEqual(expected, actual interface{}) bool {
	return !compareEquality(expected, actual)
}

// Checks that the supplied expected and actual objects are compareEquality
// this code is a copy of the ObjectsAreEqual method from :
// 		https://github.com/stretchr/testify/blob/master/assert/assertions.go
//		Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell
func compareEquality(expected, actual interface{}) bool {

	if expected == nil || actual == nil {
		return expected == actual
	}

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	if expectedValue == actualValue {
		return true
	}

	// Attempt comparison after type conversion
	if actualValue.Type().ConvertibleTo(expectedValue.Type()) && expectedValue == actualValue.Convert(expectedValue.Type()) {
		return true
	}

	// Last ditch effort
	if fmt.Sprintf("%#v", expected) == fmt.Sprintf("%#v", actual) {
		return true
	}

	return false
}

// Nil checks that the actual value is nil
func Nil(t *testing.T, actual interface{}, message ...string) {
	if !reflect.ValueOf(actual).IsNil() {
		t.Errorf("%v Expected \n[%#v]\nto be\nnil\n%v ", message, actual, callerInfo(2))
	}
}

// NotNil checks that the actual value is not nil
func NotNil(t *testing.T, actual interface{}, message ...string) {
	if reflect.ValueOf(actual).IsNil() {
		t.Errorf("%v Expected \n[%#v]\nnot to be\nnil\n%v ", message, actual, callerInfo(2))
	}
}

// True checks that the actual value is true
func True(t *testing.T, actual bool, message ...string) {
	if actual != true {
		t.Errorf("%v Expected \n[%#v]\nto be\nTrue\n%v ", message, actual, callerInfo(2))
	}
}

// False checks that the actual value is false
func False(t *testing.T, actual bool, message ...string) {
	if actual != false {
		t.Errorf("%v Expected \n[%#v]\nto be\nFalse\n%v ", message, actual, callerInfo(2))
	}
}

// Error checks that the actual error is not nil (compiler will check it supports the error interface)
func Error(t *testing.T, actual error, message ...string) {
	if actual == nil {
		t.Errorf("%v Expected \n[%#v]\nto be an error\n%v ", message, actual, callerInfo(2))
	}
}

// Error checks that the actual error is nil (compiler will check it supports the error interface)
func NotError(t *testing.T, actual error, message ...string) {
	if actual != nil {
		t.Errorf("%v Expected \n[%#v]\nto not be an error\n%v ", message, actual, callerInfo(2))
	}
}

// Checks that the lengths of the supplied Slice | Map | String are the same
func Len(t *testing.T, expected, actual interface{}, message ...string) {
	if !compareLength(expected, actual) {
		t.Errorf("%v Expected length \n[%#v]\nto be\n[%#v]\n%v ", message, expected, actual, callerInfo(2))
	}
}

// compares lengths dependant on teh type of variables passed in for comparison
// panics if the type is not Slice | Map | String
func compareLength(a, b interface{}) bool {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(a)
		return s.Len() == b
	case reflect.Map:
		m := reflect.ValueOf(a)
		return m.Len() == b
	case reflect.String:
		m := reflect.ValueOf(a).String()
		return utf8.RuneCountInString(m) == b
	}
	panic("parameter does not have a Len")
}




// callerInfo returns the file and line number from the caller stack
func callerInfo(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%v:%v", file, line)
}


