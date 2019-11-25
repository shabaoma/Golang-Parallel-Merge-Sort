package quick_sort

type Nums []int

func (a Nums) Len() int           { return len(a) }
func (a Nums) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nums) Less(i, j int) bool { return a[i] < a[j] }
