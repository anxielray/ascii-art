package colortheascii

func Less_Concat_Full(a, b, c, d, e, f, g, h []int, color, banner string) (string, string, string, string, string, string, string, string) {
	var ab, bc, cd, de, ef, fg, gh, hi string
	for _, a := range a {
		ab += ColorText(color, string(SecondaryMap(a, banner)))
	}
	for _, b := range b {
		bc += ColorText(color, string(SecondaryMap(b, banner)))
	}
	for _, c := range c {
		cd += ColorText(color, string(SecondaryMap(c, banner)))
	}
	for _, d := range d {
		de += ColorText(color, string(SecondaryMap(d, banner)))
	}
	for _, e := range e {
		ef += ColorText(color, string(SecondaryMap(e, banner)))
	}
	for _, f := range f {
		fg += ColorText(color, string(SecondaryMap(f, banner)))
	}
	for _, g := range g {
		gh += ColorText(color, string(SecondaryMap(g, banner)))
	}
	for _, h := range h {
		hi += ColorText(color, string(SecondaryMap(h, banner)))
	}
	return " " + ab + "\n", bc + "\n", cd + "\n", de + "\n", ef + "\n", fg + "\n", gh + "\n", hi
}
