package test

import (
	bubble "../pkg/sorting"
	"reflect"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	scores := []int{13, 2, 3, 5, 11, 7}

	expected := []int{2, 3, 5, 7, 11, 13}
	actual := bubble.BubbleSort(scores)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MakeSomeNoise was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}
