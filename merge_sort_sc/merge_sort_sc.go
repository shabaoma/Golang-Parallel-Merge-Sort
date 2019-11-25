package merge_sort_sc

func Sort(s []int, res chan []int) {
	if len(s) == 1 {
		res <- s
		return
	}

	mid := len(s) >> 1
	left_chan, right_chan := make(chan []int), make(chan []int)
	go Sort(s[:mid], left_chan)
	go Sort(s[mid:], right_chan)
	ls, rs := <-left_chan, <-right_chan
	close(left_chan)
	close(right_chan)
	res <- merge(ls, rs)
	return
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
