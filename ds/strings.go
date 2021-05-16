package ds

func Repeat(c string) string {
	var res string

	for i := 0; i < 5; i++ {
		res = res + c
	}

	return res
}
