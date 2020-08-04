package test

import (
	math "../pkg"
	"testing"
)

// Run this with go test
func TestAdd(t *testing.T) {
	expected := 4
	actual := math.Add(2, 2)

	if actual != expected {
		t.Errorf("Add was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}
