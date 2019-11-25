package merge_sort_sc

import (
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{1, 5, 2, 245, 574, -3, 55, -10, 0}
	ch := make(chan []int)
	go Sort(s, ch)
	res := <-ch
	t.Log(res)
}
