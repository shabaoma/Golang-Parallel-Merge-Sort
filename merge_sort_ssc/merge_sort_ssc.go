package merge_sort_ssc

func Sort(s []int) []int {
	l_ch, r_ch, res := make(chan []int), make(chan []int), make(chan []int)
	quarter := len(s) >> 2
	mid := len(s) >> 1
	t_quarter := quarter + mid

	go func() {
		l_ch <- goSort(s[:quarter])
	}()
	go func() {
		l_ch <- goSort(s[quarter:mid])
	}()
	go func() {
		r_ch <- goSort(s[mid:t_quarter])
	}()
	go func() {
		r_ch <- goSort(s[t_quarter:])
	}()

	go func() {
		res <- merge(<-l_ch, <-l_ch)
	}()
	go func() {
		res <- merge(<-r_ch, <-r_ch)
	}()

	return merge(<-res, <-res)
}

func goSort(s []int) []int {
	if len(s) == 1 {
		return s
	}
	mid := len(s) >> 1
	return merge(goSort(s[:mid]), goSort(s[mid:]))
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
