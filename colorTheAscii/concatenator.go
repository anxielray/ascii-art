package colortheascii

func Concatenator(Index, A, B, C, D, E, F, G, H []int, color, letters string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	indices := addLetterLen(Index, letters)
	for i, a := range A {
		if ContainIndex(indices, i) {
			ab += ColorText(color, Maps(A[i]))
		} else {
			ab += Maps(a)
		}
	}
	for i, b := range B {
		if ContainIndex(indices, i) {
			bc += ColorText(color, Maps(B[i]))
		} else {
			bc += Maps(b)
		}
	}
	for i, c := range C {
		if ContainIndex(indices, i) {
			cd += ColorText(color, Maps(C[i]))
		} else {
			cd += Maps(c)
		}
	}
	for i, d := range D {
		if ContainIndex(indices, i) {
			de += ColorText(color, Maps(D[i]))
		} else {
			de += Maps(d)
		}
	}
	for i, e := range E {
		if ContainIndex(indices, i) {
			ef += ColorText(color, Maps(E[i]))
		} else {
			ef += Maps(e)
		}
	}
	for i, f := range F {
		if ContainIndex(indices, i) {
			fg += ColorText(color, Maps(F[i]))
		} else {
			fg += Maps(f)
		}
	}
	for i, g := range G {
		if ContainIndex(indices, i) {
			gh += ColorText(color, Maps(G[i]))
		} else {
			gh += Maps(g)
		}
	}
	for i, h := range H {
		if ContainIndex(indices, i) {
			hi += ColorText(color, Maps(H[i]))
		} else {
			hi += Maps(h)
		}
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}

func ContainIndex(a []int, num int) bool {
	for _, nums := range a {
		if nums == num {
			return true
		}
	}
	return false
}

func addLetterLen(a []int, letters string) (added []int) {
	for _, nums := range a {
		for x := 0; x < len(letters); x++ {
			added = append(added, nums+x)
		}
	}
	return
}
