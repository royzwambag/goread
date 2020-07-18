package goread

import (
	"reflect"
	"testing"
)

func TestStringSliceToIntSlice(t *testing.T) {
	input := []string{"1", "2", "3", "4", "5"}
	expectedResult := []int{1, 2, 3, 4, 5}

	result := stringSliceToIntSlice(input)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Got wrong output, expected: %v, got: %v", expectedResult, result)
	}
}

func BenchmarkStringSliceToIntSlice(b *testing.B) {
	strings := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for i := 0; i < b.N; i++ {
		stringSliceToIntSlice(strings)
	}
}
