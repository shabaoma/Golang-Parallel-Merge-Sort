package merge_sort_s

func Sort(s []int) []int {
	if len(s) == 1 {
		return s
	}
	mid := len(s) >> 1
	return merge(Sort(s[:mid]), Sort(s[mid:]))
}

func merge(s1, s2 []int) []int {
	s := make([]int, len(s1)+len(s2))
	i, j, k := 0, 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			s[k] = s1[i]
			i++
		} else {
			s[k] = s2[j]
			j++
		}
		k++
	}
	for i < len(s1) {
		s[k] = s1[i]
		i++
		k++
	}
	for j < len(s2) {
		s[k] = s2[j]
		j++
		k++
	}
	return s
}
