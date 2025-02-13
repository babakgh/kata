package grokking_test

import (
	"reflect"
	"testing"
)

func assertEqual[K comparable](t *testing.T, got, want K) {
	t.Helper()
	if got != want {
		t.Errorf("incorrect, got: %v, want %v", got, want)
	}
}

func assertDeepEqual[T any](t *testing.T, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
