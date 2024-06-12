package colortheascii

func ColorConcatenator(Indx, A, B, C, D, E, F, G, H []int, color string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	for _, indx := range Indx {
		for i, a := range A {
			if i == indx {
				ab += SecondaryMap(A[indx], color)
			} else {
				ab += Maps(a)
			}
		}
	}

	for _, indx := range Indx {
		for i, b := range B {
			if i == indx {
				bc += SecondaryMap(B[indx], color)
			} else {
				bc += Maps(b)
			}
		}
	}

	for _, indx := range Indx {
		for i, c := range C {
			if i == indx {
				cd += SecondaryMap(c, color)
			} else {
				cd += Maps(c)
			}
		}
	}

	for _, indx := range Indx {
		for i, d := range D {
			if i == indx {
				de += SecondaryMap(d, color)
			} else {
				de += Maps(d)
			}
		}
	}

	for _, indx := range Indx {
		for i, e := range E {
			if i == indx {
				ef += SecondaryMap(E[indx], color)
			} else {
				ef += Maps(e)
			}
		}
	}

	for _, indx := range Indx {
		for i, f := range F {
			if i == indx {
				fg += SecondaryMap(F[indx], color)
			} else {
				fg += Maps(f)
			}
		}
	}

	for _, indx := range Indx {
		for i, g := range G {
			if i == indx {
				gh += SecondaryMap(G[indx], color)
			} else {
				gh += Maps(g)
			}
		}
	}

	for _, indx := range Indx {
		for i, h := range H {
			if i == indx {
				hi += SecondaryMap(H[indx], color)
			} else {
				hi += Maps(h)
			}
		}
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}
