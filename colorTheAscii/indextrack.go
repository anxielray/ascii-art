package colortheascii

func IndexTracker(s, s1 string) (Index []int) {
	wanted, n, exact := 0, len(s), 0
	for i := range s1 {
		if s[0] == s1[i] {
			wanted = i
			if s1[wanted:(wanted+n)] == s {
				exact = wanted
				Index = append(Index, exact)
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return
}
