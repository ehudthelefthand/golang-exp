package list_test

import (
	"exp/list"
	"testing"
)

func TestRemove(t *testing.T) {
	// Arrange
	numbers := []int{1, 2, 3}
	// Act
	result, err := list.Remove(numbers, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Assert
	if len(result) != 2 {
		t.Errorf("expect len(result) = 3 but got %v", len(result))
	}
	if result[0] == 1 {
		t.Errorf("expect result[0] != 1 but got %v", result[0])
	}

	result, err = list.Remove(numbers, -1)
	if err == nil {
		t.Errorf("expected error but got nil")
	}
	// Reset
}
