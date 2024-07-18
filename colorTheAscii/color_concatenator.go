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
func Concatenator(Index, A, B, C, D, E, F, G, H []int, color, letters string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	indices := addLetterLen(Index, letters)
	for i, a := range A {
		if containIndex(indices, i) {
			ab += ColorText(color, Maps(A[i]))
		} else {
			ab += Maps(a)
		}
	}
	for i, b := range B {
		if containIndex(indices, i) {
			bc += ColorText(color, Maps(B[i]))
		} else {
			bc += Maps(b)
		}
	}
	for i, c := range C {
		if containIndex(indices, i) {
			cd += ColorText(color, Maps(C[i]))
		} else {
			cd += Maps(c)
		}
	}
	for i, d := range D {
		if containIndex(indices, i) {
			de += ColorText(color, Maps(D[i]))
		} else {
			de += Maps(d)
		}
	}
	for i, e := range E {
		if containIndex(indices, i) {
			ef += ColorText(color, Maps(E[i]))
		} else {
			ef += Maps(e)
		}
	}
	for i, f := range F {
		if containIndex(indices, i) {
			fg += ColorText(color, Maps(F[i]))
		} else {
			fg += Maps(f)
		}
	}
	for i, g := range G {
		if containIndex(indices, i) {
			gh += ColorText(color, Maps(G[i]))
		} else {
			gh += Maps(g)
		}
	}
	for i, h := range H {
		if containIndex(indices, i) {
			hi += ColorText(color, Maps(H[i]))
		} else {
			hi += Maps(h)
		}
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}

func containIndex(a []int, num int) bool {
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
