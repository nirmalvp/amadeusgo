package params

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWith(t *testing.T) {
	testCases := []struct {
		key      string
		value    interface{}
		expected Params
	}{
		{"key", "value", Params{"key": "value"}},
		{"key", 2, Params{"key": fmt.Sprintf("%d", 2)}},
		{"key", 2.54, Params{"key": fmt.Sprintf("%f", 2.54)}},
	}

	for _, testCase := range testCases {
		gotParam := With(testCase.key, testCase.value)
		if !reflect.DeepEqual(gotParam, testCase.expected) {
			t.Errorf("TestWith, got: %v, want: %v.", gotParam, testCase.expected)
		}
	}
}

func TestAnd(t *testing.T) {
	param := With("with_key", "with_value")
	testCases := []struct {
		key      string
		value    interface{}
		expected Params
	}{
		{"key", "value", Params{"with_key": "with_value", "key": "value"}},
		{"key", 2, Params{"with_key": "with_value", "key": fmt.Sprintf("%d", 2)}},
		{"key", 2.54, Params{"with_key": "with_value", "key": fmt.Sprintf("%f", 2.54)}},
	}

	for _, testCase := range testCases {
		gotParam := param.And(testCase.key, testCase.value)
		if !reflect.DeepEqual(gotParam, testCase.expected) {
			t.Errorf("TestWith, got: %v, want: %v.", gotParam, testCase.expected)
		}
	}
}
