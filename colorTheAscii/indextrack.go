package colortheascii

func IndexTracker(s, s1 string) (focusIndex []int) {
	sliceS := []rune(s)
	for _, char := range sliceS {
		for j := range s1 {
			if char == rune(s1[j]) {
				focusIndex = append(focusIndex, j)
				continue
			}
		}
	}
	return
}
