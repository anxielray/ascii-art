package colortheascii

func Concatenator(indx int, A, B, C, D, E, F, G, H []int, color, letters string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	for i := range A {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				ab += string(SecondaryMap(A[x], color))
			}
		} else {
			ab += Maps(A[i])
		}
	}
	for i := range B {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				bc += string(SecondaryMap(B[x], color))
			}
		} else {
			bc += Maps(B[i])
		}
	}
	for i := range C {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				cd += string(SecondaryMap(C[x], color))
			}
		} else {
			cd += Maps(C[i])
		}
	}
	for i := range D {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				de += string(SecondaryMap(D[x], color))
			}
		} else {
			de += Maps(D[i])
		}
	}
	for i := range E {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				ef += string(SecondaryMap(E[x], color))
			}
		} else {
			ef += Maps(E[i])
		}
	}
	for i := range F {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				fg += string(SecondaryMap(F[x], color))
			}
		} else {
			fg += Maps(F[i])
		}
	}
	for i := range G {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				gh += string(SecondaryMap(G[x], color))
			}
		} else {
			gh += Maps(G[i])
		}
	}
	for i := range H {
		if i == indx {
			for x := i; x <= (len(letters)+i)-1; x++ {
				hi += string(SecondaryMap(H[x], color))
			}
		} else {
			hi += Maps(H[i])
		}
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}
