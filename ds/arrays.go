package ds

func SumList(l []int) int {
	s := 0

	for _, num := range l {
		s += num
	}

	return s
}

func SumAllList(lists ...[]int) []int {
	var s int = 0
	var res []int //:= make([]int, len(lists))

	for _, l := range lists {
		s = SumList(l)
		res = append(res, s)
	}

	return res
}
