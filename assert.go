package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"unicode/utf8"
)

// checkFn provides the prototype for a function that takes two parameters and returns a bool to indicate a successful or failed assertion.
type checkFn func(a, b interface{}) bool

// Assert is the main assertion function it takes
// 	a reference to the testing.T instance,
// 	a value to check,
// 	a function to be applied,
// 	the expected value,
// 	a message to be displayed on assertion failure
func Assert(t *testing.T, value interface{}, check checkFn, expected interface{}, message string) bool {
	result := check(value, expected)
	if !result {
		t.Errorf("ASSERT: [%v] Expected \n[%#v]\nto be\n[%#v]\n%v ", message, value, expected, callerInfo(2))
	}
	return !result
}

// caller info returns the file and line number from the caller stack
func callerInfo(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%v:%v", file, line)
}

// EqualDeep compares deeply (using reflect.deep) the value and expected value to see if they are equal
func EqualDeep(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// NotEqualDeep compares deeply (using reflect.deep) the value and expected value to see if they are NOT equal
func NotEqualDeep(a, b interface{}) bool {
	return !reflect.DeepEqual(a, b)
}

// EqualString compares as strings the value and expected value to see if they are equal
func EqualString(a, b interface{}) bool {
	return reflect.ValueOf(a).String() == reflect.ValueOf(b).String()
}

// EqualInt compares as integers the value and expected value to see if they are equal
func EqualInt(a, b interface{}) bool {
	return reflect.ValueOf(a).Int() == reflect.ValueOf(b).Int()
}

// NotEqualInt compares as integers the value and expected value to see if they are NOT equal
func NotEqualInt(a, b interface{}) bool {
	return !(reflect.ValueOf(a).Int() == reflect.ValueOf(b).Int())
}

// EqualFloat compares as floats the value and expected value to see if they are equal
func EqualFloat(a, b interface{}) bool {
	return reflect.ValueOf(a).Float() == reflect.ValueOf(b).Float()
}

// NotEqualFloat compares as floats the value and expected value to see if they are NOT equal
func NotEqualFloat(a, b interface{}) bool {
	return !(reflect.ValueOf(a).Float() == reflect.ValueOf(b).Float())
}

func Len(a, b interface{}) bool {
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
	panic("Value parameter not an Array")
}


// Checks that the supplied expected and actual objects are equal
// this code is a copy of the ObjectsAreEqual method from :
// 		https://github.com/stretchr/testify/blob/master/assert/assertions.go
//		Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell
func Equal(expected, actual interface{}) bool {

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

func NotEqual(expected, actual interface{}) bool {
	return !Equal(expected, actual)
}