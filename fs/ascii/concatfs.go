package ascii

func Concatenatorfs(a, b, c, d, e, f, g, h []int) (str1, str2, str3, str4, str5, str6, str7, str8 string) {
	for _, num := range a {
		str1 += Mapfs(num)
	}
	for _, num := range b {
		str2 += Mapfs(num)
	}
	for _, num := range c {
		str3 += Mapfs(num)
	}
	for _, num := range d {
		str4 += Mapfs(num)
	}
	for _, num := range e {
		str5 += Mapfs(num)
	}
	for _, num := range f {
		str6 += Mapfs(num)
	}
	for _, num := range g {
		str7 += Mapfs(num)
	}
	for _, num := range h {
		str8 += Mapfs(num)
	}
	return " " + str1 + "\n", str2 + "\n", str3 + "\n", str4 + "\n", str5 + "\n", str6 + "\n", str7 + "\n", str8
}
