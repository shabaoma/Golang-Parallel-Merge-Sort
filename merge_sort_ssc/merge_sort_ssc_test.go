package merge_sort_ssc

import (
	"testing"
)

func TestSort(t *testing.T) {
	s := []int{1, 5, 2, 245, 574, -3, 55, -10, 0}
	res := Sort(s)
	t.Log(res)
}
