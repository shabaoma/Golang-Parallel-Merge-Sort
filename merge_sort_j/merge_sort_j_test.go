package merge_sort_j

import (
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{1, 5, 2, 245, 574, -3, 55, -10, 0}
	ch := make(chan int, len(s))
	Sort(s, ch)
	var res []int
	for v := range ch {
		res = append(res, v)
	}
	t.Log(res)
}
