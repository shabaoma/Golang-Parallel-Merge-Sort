package main

import (
	"fmt"
	"sort"
	"sort_test/merge_sort_j"
	"sort_test/merge_sort_s"
	// "sort_test/merge_sort_sc"
	"sort_test/merge_sort_ssc"
	"sort_test/quick_sort"
	"testing"
)

var slice_length = 10000000

func TestQuickSort(t *testing.T) {
	s := quick_sort.Nums(GenerateSlice(slice_length))
	sort.Sort(s)
}

func TestMergeSortS(t *testing.T) {
	s := GenerateSlice(slice_length)
	_ = merge_sort_s.Sort(s)
}

// func TestMergeSortSC(t *testing.T) {
// 	s := GenerateSlice(slice_length)
// 	ch := make(chan []int)
// 	merge_sort_sc.Sort(s, ch)
// }

func TestMergeSortSSC(t *testing.T) {
	s := GenerateSlice(slice_length)
	_ = merge_sort_ssc.Sort(s)
}

func TestMergeSortJ(t *testing.T) {
	s := GenerateSlice(slice_length)
	ch := make(chan int, len(s))
	merge_sort_j.Sort(s, ch)
}

// func BenchmarkQuickSort(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := quick_sort.Nums(GenerateSlice(slice_length))
// 		sort.Sort(s)
// 	}
// }

// func BenchmarkMergeSortS(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := GenerateSlice(slice_length)
// 		_ = merge_sort_s.Sort(s)
// 	}
// }

// func BenchmarkMergeSortSC(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := GenerateSlice(slice_length)
// 		ch := make(chan []int)
// 		merge_sort_sc.Sort(s, ch)
// 	}
// }

// func BenchmarkMergeSortSSC(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := GenerateSlice(slice_length)
// 		_ = merge_sort_ssc.Sort(s)
// 	}
// }

// func BenchmarkMergeSortJ(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := GenerateSlice(slice_length)
// 		ch := make(chan int, len(s))
// 		merge_sort_j.Sort(s, ch)
// 	}
// }

func BenchmarkQuickSortIncrease(b *testing.B) {
	for n := 10; n <= slice_length; n *= 10 {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := quick_sort.Nums(GenerateSlice(n))
				sort.Sort(s)
			}
		})
	}
}

func BenchmarkMergeSortSIncrease(b *testing.B) {
	for n := 10; n <= slice_length; n *= 10 {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := GenerateSlice(n)
				_ = merge_sort_s.Sort(s)
			}
		})
	}
}

// func BenchmarkMergeSortSCIncrease(b *testing.B) {
// 	for n := 10; n <= slice_length; n *= 10 {
// 		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				s := GenerateSlice(n)
// 				ch := make(chan []int)
// 				merge_sort_sc.Sort(s, ch)
// 			}
// 		})
// 	}
// }

func BenchmarkMergeSortSSCIncrease(b *testing.B) {
	for n := 10; n <= slice_length; n *= 10 {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := GenerateSlice(n)
				_ = merge_sort_ssc.Sort(s)
			}
		})
	}
}

func BenchmarkMergeSortJIncrease(b *testing.B) {
	for n := 10; n <= slice_length; n *= 10 {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := GenerateSlice(n)
				ch := make(chan int, len(s))
				merge_sort_j.Sort(s, ch)
			}
		})
	}
}
