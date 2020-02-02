package main

import "testing"

type QuickSortCase struct {
	input    []int
	expected []int
}

var quickSortCases = []QuickSortCase{
	{[]int{1}, []int{1}},
	{[]int{2, 1}, []int{1, 2}},
	{[]int{1, 3, 2, 1}, []int{1, 1, 2, 3}},
}

func TestQuickSort(t *testing.T) {
	for k, c := range quickSortCases {
		result := QuickSort(c.input)

		if !Compare(result, c.expected) {
			t.Error("TestQuickSort: failed case №", k, "Input:", c.input, "Result:", result, "Expected", c.expected)
		}
	}
}

type CompareCase struct {
	a        []int
	b        []int
	expected bool
}

var compareCases = []CompareCase{
	{[]int{1}, []int{1}, true},
	{[]int{1}, []int{2}, false},
	{[]int{1, 1}, []int{2}, false},
	{[]int{1, 2}, []int{1, 2}, true},
	{[]int{2, 1}, []int{3, 1}, false},
}

func TestCompare(t *testing.T) {
	for k, c := range compareCases {
		if Compare(c.a, c.b) != c.expected {
			t.Error("TestCompare: провален кейс №", k, c.a, c.b, "Expected", c.expected)
		}
	}
}

// 1000,     // тысяча
// 10000,    // десять тысяч
// 100000,   // сто тысяч
// 1000000,  // миллион
// 10000000, // десять миллионов

func BenchmarkQuickSort1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortFile("./nums_1000")
	}
}

func BenchmarkQuickSort10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortFile("./nums_10000")
	}
}

func BenchmarkQuickSort100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortFile("./nums_100000")
	}
}

func BenchmarkQuickSort1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortFile("./nums_1000000")
	}
}

func BenchmarkQuickSort10000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortFile("./nums_10000000")
	}
}
