package array

import (
	"testing"
)

func TestUnion(t *testing.T) {
	arr1 := []string{"1", "2", "3", "4"}
	arr2 := []string{"3", "4", "5", "6"}
	want := []string{"1", "2", "3", "4", "5", "6"}
	got := Union(arr1, arr2)
	t.Logf("got: %v, want: %v", got, want)
}

func TestIntersection(t *testing.T) {
	arr1 := []string{"1", "2", "3", "4"}
	arr2 := []string{"3", "4", "5", "6"}
	want := []string{"3", "4"}
	got := Intersection(arr1, arr2)
	t.Logf("got: %v, want: %v", got, want)
}

func TestDifference(t *testing.T) {
	arr1 := []string{"1", "2", "3", "4"}
	arr2 := []string{"3", "4", "5", "6"}
	want := []string{"1", "2", "5", "6"}
	got := Difference(arr1, arr2)
	t.Logf("got: %v, want: %v", got, want)
}
