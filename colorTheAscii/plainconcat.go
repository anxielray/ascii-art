package colortheascii

func PlainConcatenator(a, b, c, d, e, f, g, h []int, color string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	for _, a := range a {
		ab += ColorText(color, string(Maps(a)))
	}
	for _, b := range b {
		bc += ColorText(color, string(Maps(b)))
	}
	for _, c := range c {
		cd += ColorText(color, string(Maps(c)))
	}
	for _, d := range d {
		de += ColorText(color, string(Maps(d)))
	}
	for _, e := range e {
		ef += ColorText(color, string(Maps(e)))
	}
	for _, f := range f {
		fg += ColorText(color, string(Maps(f)))
	}
	for _, g := range g {
		gh += ColorText(color, string(Maps(g)))
	}
	for _, h := range h {
		hi += ColorText(color, string(Maps(h)))
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}
