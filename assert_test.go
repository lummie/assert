package assert

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

type StringType string

const (
	STRING_CONST      string     = "string_const"
	STRING_CONST_TYPE StringType = "string_const_type"
)

func TestDeepEqual(t *testing.T) {
	if !reflect.DeepEqual(1, 1) {
		t.Error("Should be equal")
	}
	if !reflect.DeepEqual("", "") {
		t.Error("Should be equal")
	}
	if reflect.DeepEqual(int64(1), int32(1)) {
		t.Error("should be different")
	}
	if reflect.DeepEqual(int16(1), int32(1)) {
		t.Error("Should be different")
	}
	if !reflect.DeepEqual(0.0123, 0.0123) {
		t.Error("Should be Equal")
	}
	if reflect.DeepEqual(0.0123, 0.01234) {
		t.Error("Should be Different")
	}
	if !reflect.DeepEqual("abcdef", "abcdef") {
		t.Error("Should be Equal")
	}
	if !reflect.DeepEqual("string_const", STRING_CONST) {
		t.Error("Should be Equal")
	}
	if reflect.DeepEqual("string_const_type", STRING_CONST_TYPE) {
		t.Error("Should be different")
	}
}



func TestLen(t *testing.T) {
	// arrays
	a := make([]string, 1, 1)
	if !compareLength(a, 1) {
		t.Error("Array Length should be 1")
	}
	a2 := make([]string, 0, 1)
	if !compareLength(a2, 0) {
		t.Error("Array Length should be 0")
	}

	//maps
	m1 := make(map[int]string)
	if !compareLength(m1, 0) {
		t.Error("Map Length should be 0")
	}
	m2 := make(map[int]string)
	m2[1] = "fred"
	if !compareLength(m2, 1) {
		t.Error("Map Length should be 1")
	}
	m3 := make(map[int]string)
	m3[1] = "fred"
	m3[2] = "fred"
	if !compareLength(m3, 2) {
		t.Error("Map Length should be 2")
	}

	//strings
	if !compareLength("fred", 4) {
		t.Error("String Length should be 4")
	}
	if !compareLength("", 0) {
		t.Error("String Length should be 4")
	}
	if !compareLength("億種の商品をい", 7) {
		t.Errorf("String Length should be %d", utf8.RuneCountInString("億種の商品をい"))
	}

}
