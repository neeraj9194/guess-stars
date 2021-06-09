package main

import (
	"testing"
)

// If an element exists in a slice
func TestElementExist(t *testing.T) {
	l := []int{1, 2, 4, 5}
	if !ElementExist(l, 5) {
		t.Fatal("Failed.")
	}
}


func TestElementExistEmptyList(t *testing.T) {
	l := []int{}
	if ElementExist(l, 5) {
		t.Fatal("Failed.")
	}
}
