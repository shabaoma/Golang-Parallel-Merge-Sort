package merge_sort_s

import (
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{1, 5, 2, 245, 574, -3, 55, -10, 0}
	res := Sort(s)
	t.Log(res)
}

func TestMerge(t *testing.T) {
	s1 := []int{1, 3, 5, 7, 9, 33, 56, 78}
	s2 := []int{2, 4, 6, 8, 14, 34, 45, 66}
	res := merge(s1, s2)
	t.Log(res)
}
