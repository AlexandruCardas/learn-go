package test

import (
	sentences "../pkg"
	"testing"
)

// Run this with go test
func TestMakeSomeNoise(t *testing.T) {
	expected := "COOL!"
	actual := sentences.MakeSomeNoise("cool")

	if actual != expected {
		t.Errorf("MakeSomeNoise was incorrect! Expected: %s, Actual: %s", expected, actual)
	}
}
