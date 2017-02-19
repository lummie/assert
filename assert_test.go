package assert_test

import (
	"github.com/lummie/assert"
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

func TestEqualString(t *testing.T) {
	if !assert.EqualString("string_const", STRING_CONST) {
		t.Error("Should be Equal")
	}
	if !assert.EqualString("string_const_type", STRING_CONST_TYPE) {
		t.Error("Should be Equal")
	}
	if assert.EqualString("1", 1) {
		t.Error("Should be different")
	}
	if assert.EqualString("1.0", 1.0) {
		t.Error("Should be different")
	}
}

func TestEqualDeep(t *testing.T) {
	if !assert.EqualDeep(1, 1) {
		t.Error("Should be equal")
	}
	if assert.EqualDeep(1, 2) != false {
		t.Error("Should be different")
	}
}

func TestNotEqualDeep(t *testing.T) {
	if !assert.NotEqualDeep(1, 2) {
		t.Error("Should not be equal")
	}
	if assert.NotEqualDeep(1, 1) {
		t.Error("Should not be different")
	}
	if !assert.NotEqualDeep(int64(1), int32(2)) {
		t.Error("Should not be equal")
	}
}

func TestEqualInt(t *testing.T) {
	if !assert.EqualInt(int32(1), int32(1)) {
		t.Error("Should be equal")
	}
	if !assert.EqualInt(int32(-10), int32(-10)) {
		t.Error("Should be equal")
	}
	if !assert.EqualInt(int64(1), int32(1)) {
		t.Error("Should be equal")
	}
	if !assert.EqualInt(int64(-1), int32(-1)) {
		t.Error("Should be equal")
	}
	if !assert.EqualInt(int64(2^^31), int32(2^^31)) {
		t.Error("Should be equal")
	}
	if assert.EqualInt(int64(2^^32), int32(2^^31)) {
		t.Error("Should be different")
	}
}

func TestNotEqualInt(t *testing.T) {
	if assert.NotEqualInt(int32(1), int32(1)) {
		t.Error("Should be equal")
	}
	if assert.NotEqualInt(int32(-10), int32(-10)) {
		t.Error("Should be equal")
	}
	if assert.NotEqualInt(int64(1), int32(1)) {
		t.Error("Should be equal")
	}
	if assert.NotEqualInt(int64(-1), int32(-1)) {
		t.Error("Should be equal")
	}
	if assert.NotEqualInt(int64(2^^31), int32(2^^31)) {
		t.Error("Should be equal")
	}
	if !assert.NotEqualInt(int64(2^^32), int32(2^^31)) {
		t.Error("Should be different")
	}
}

func TestLen(t *testing.T) {
	// arrays
	a := make([]string, 1, 1)
	if !assert.Len(a, 1) {
		t.Error("Array Length should be 1")
	}
	a2 := make([]string, 0, 1)
	if !assert.Len(a2, 0) {
		t.Error("Array Length should be 0")
	}

	//maps
	m1 := make(map[int]string)
	if !assert.Len(m1, 0) {
		t.Error("Map Length should be 0")
	}
	m2 := make(map[int]string)
	m2[1] = "fred"
	if !assert.Len(m2, 1) {
		t.Error("Map Length should be 1")
	}
	m3 := make(map[int]string)
	m3[1] = "fred"
	m3[2] = "fred"
	if !assert.Len(m3, 2) {
		t.Error("Map Length should be 2")
	}

	//strings
	if !assert.Len("fred", 4) {
		t.Error("String Length should be 4")
	}
	if !assert.Len("", 0) {
		t.Error("String Length should be 4")
	}
	if !assert.Len("億種の商品をい", 7) {
		t.Errorf("String Length should be %d", utf8.RuneCountInString("億種の商品をい"))
	}

}
