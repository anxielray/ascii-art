package colortheascii

func IndexTracker(substring string, str string) (result []int) {
	result = make([]int, 0)
	for i := 0; i <= len(str); i++ {
		for j := i; j <= len(str); j++ {
			if str[i:j] == substring {
				result = append(result, i)
			}
		}
	}
	return
}
