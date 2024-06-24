package colortheascii

func IndexTracker(s, s1 string) (Index int) {
	wanted, n := 0, len(s)
	for i := range s1 {
		if s[0] == s1[i] {
			wanted = i
			if s1[wanted:(wanted+n)] == s {
				Index = wanted
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return
}
