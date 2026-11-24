package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	res := generateRandomElements(10)
	if len(res) != 10 {
		t.Error("Expected 10 elements, got ", len(res))
	}
	res_1 := generateRandomElements(0)
	if len(res_1) != 0 {
		t.Error("Expected 0 elements, got ", len(res_1))
	}
	res_2 := generateRandomElements(1)
	if len(res_2) != 1 {
		t.Error("Expected 1 element, got", len(res))
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty", []int{}, 0},
		{"one", []int{1}, 1},
		{"two", []int{1, 2}, 2},
		{"minus", []int{-5, -1, -10}, -1},
		{"plus", []int{5, 5, 5}, 5},
		{"mixed", []int{-5, 0, 3, 2}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	data := []int{1, 50, 100, 2, 4, 8, 9, 10, 500, 3, 20, 30, 40, 55, 66, 77}
	expected := 500
	res := maxChunks(data)
	if res != expected {
		t.Error("Expected", expected, "got", res)
	}
	if res := maxChunks([]int{}); res != 0 {
		t.Error("Expected 0, got", res)
	}

	if res := maxChunks([]int{42}); res != 42 {
		t.Error("Expected 42, got", res)
	}

	if res := maxChunks([]int{1, 2, 3, 4, 5}); res != 5 {
		t.Error("Expected 5, got", res)
	}

	if res := maxChunks([]int{-5, -10, -1}); res != -1 {
		t.Error("Expected -1, got", res)
	}
}
