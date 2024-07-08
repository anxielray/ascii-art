package colortheascii

func Concatenator(Index, A, B, C, D, E, F, G, H []int, color, letters string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	for i, a := range A {
		if !(ContainIndex(Index, i, letters)) {
			ab += Maps(a)
		} else {
			ab += ColorText(color, Maps(a))
		}
	}
	for i, b := range B {
		if !(ContainIndex(Index, i, letters)) {
			bc += Maps(b)
		} else {
			bc += ColorText(color, Maps(b))
		}
	}
	for i, c := range C {
		if !(ContainIndex(Index, i, letters)) {
			cd += Maps(c)
		} else {
			cd += ColorText(color, Maps(c))
		}
	}
	for i, d := range D {
		if !(ContainIndex(Index, i, letters)) {
			de += Maps(d)
		} else {
			de += ColorText(color, Maps(d))
		}
	}
	for i, e := range E {
		if !(ContainIndex(Index, i, letters)) {
			ef += Maps(e)
		} else {
			ef += ColorText(color, Maps(e))
		}
	}
	for i, f := range F {
		if !(ContainIndex(Index, i, letters)) {
			fg += Maps(f)
		} else {
			fg += ColorText(color, Maps(f))
		}
	}
	for i, g := range G {
		if !(ContainIndex(Index, i, letters)) {
			gh += Maps(g)
		} else {
			gh += ColorText(color, Maps(g))
		}
	}
	for i, h := range H {
		if !(ContainIndex(Index, i, letters)) {
			hi += Maps(h)
		} else {
			hi += ColorText(color, Maps(h))
		}
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}

func ContainIndex(a []int, num int, letters string) bool {
	for _, nums := range a {
		for x := nums; x < len(letters); x++ {
			if x == num {
				return true
			}
		}
	}
	return false
}
