package colortheascii

/*
The function takes  slices of integers and calls the Map(int)string function
which opens the banner file, and concatenates the lines together to form a
single line of the ASCII art.The function also takes another slice called Index,
from the IndexTracker(substr, str)[]int, which contains the starting index of the substring letters to be colored.
The function further calls other sub functoins containIndex([]int, int)bool
which will confirm if a certain index passed is contained in the slice.
It also calls for addLetterLen([]int, letters)newSlice this slice takes
a slice and the substring letters, and increaments every index that comes in between
the index to the length of the sub-string. The resulting slice will be longer.
*/
func Full_concatenator(Index, A, B, C, D, E, F, G, H []int, color, letters, banner string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	indices := addLetterLen(Index, letters)
	for i, a := range A {
		if containIndex(indices, i) {
			ab += ColorText(color, SecondaryMap(A[i], banner))
		} else {
			ab += SecondaryMap(a, banner)
		}
	}
	for i, b := range B {
		if containIndex(indices, i) {
			bc += ColorText(color, SecondaryMap(B[i], banner))
		} else {
			bc += SecondaryMap(b, banner)
		}
	}
	for i, c := range C {
		if containIndex(indices, i) {
			cd += ColorText(color, SecondaryMap(C[i], banner))
		} else {
			cd += SecondaryMap(c, banner)
		}
	}
	for i, d := range D {
		if containIndex(indices, i) {
			de += ColorText(color, SecondaryMap(D[i], banner))
		} else {
			de += SecondaryMap(d, banner)
		}
	}
	for i, e := range E {
		if containIndex(indices, i) {
			ef += ColorText(color, SecondaryMap(E[i], banner))
		} else {
			ef += SecondaryMap(e, banner)
		}
	}
	for i, f := range F {
		if containIndex(indices, i) {
			fg += ColorText(color, SecondaryMap(F[i], banner))
		} else {
			fg += SecondaryMap(f, banner)
		}
	}
	for i, g := range G {
		if containIndex(indices, i) {
			gh += ColorText(color, SecondaryMap(G[i], banner))
		} else {
			gh += SecondaryMap(g, banner)
		}
	}
	for i, h := range H {
		if containIndex(indices, i) {
			hi += ColorText(color, SecondaryMap(H[i], banner))
		} else {
			hi += SecondaryMap(h, banner)
		}
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}
