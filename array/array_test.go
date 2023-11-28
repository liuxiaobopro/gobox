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
	arr1 := []string{"1", "2", "3", "4", "4"}
	arr2 := []string{"3", "4", "5", "6"}
	want := []string{"1", "2", "5", "6"}
	got := Difference(arr1, arr2)
	t.Logf("got: %v, want: %v", got, want)
}

func TestIsIn(t *testing.T) {
	arr1 := []string{"1", "2", "3", "4"}
	arr2 := "1"
	want := true
	got := IsIn(arr1, arr2)
	t.Logf("got: %v, want: %v", got, want)

	arr11 := []string{"1", "2", "3", "4"}
	arr21 := "5"
	want1 := false
	got1 := IsIn(arr11, arr21)
	t.Logf("got1: %v, want1: %v", got1, want1)
}

func TestDelete(t *testing.T) {
	arr1 := []string{"1", "2", "3", "4"}
	want := []string{"3", "4"}
	got := Delete(arr1, "1", "2")
	t.Logf("got: %v, want: %v", got, want)
}

func TestDeleteByIndex(t *testing.T) {
	arr1 := []string{"1", "2", "3", "4"}
	want := []string{"1", "3", "4"}
	got := DeleteByIndex(arr1, 1)
	t.Logf("got: %v, want: %v", got, want)
}
