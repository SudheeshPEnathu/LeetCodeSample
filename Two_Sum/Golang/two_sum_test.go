package main

import (
	"reflect"
	"testing"
)

func TestTwoSumS1(t *testing.T) {
	t.Log("Testing TestTwoSumS1")
	got := twoSumS1([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 1")
	}

	got = twoSumS1([]int{3, 2, 4}, 6)
	want = []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 2")
	}

	got = twoSumS1([]int{2, 7, 11, 15}, 5)
	want = nil
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 3")
	}
}

func TestTwoSumS2(t *testing.T) {
	t.Log("Testing TestTwoSumS1")
	got := twoSumS2([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 1")
	}

	got = twoSumS2([]int{3, 2, 4}, 6)
	want = []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 2")
	}

	got = twoSumS2([]int{2, 7, 11, 15}, 5)
	want = nil
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 3")
	}
}

func TestTwoSumS3(t *testing.T) {
	t.Log("Testing TestTwoSumS1")
	got := twoSumS3([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 1")
	}

	got = twoSumS3([]int{3, 2, 4}, 6)
	want = []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 2")
	}

	got = twoSumS3([]int{2, 7, 11, 15}, 5)
	want = nil
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	} else {
		t.Log("Testing success for case 3")
	}
}

func BenchmarkTwoSumS1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumS1([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumS2([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumS3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumS3([]int{2, 7, 11, 15}, 9)
	}
}
