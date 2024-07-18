package colortheascii

/*
The function takes  slices of integers and calls the Map(int)string function
which opens the banner file, and concatenates the lines together to form a
single line of the ASCII art.
*/
func Fs_Concatenator(a, b, c, d, e, f, g, h []int, banner string) (str1, str2, str3, str4, str5, str6, str7, str8 string) {
	for _, num := range a {
		str1 += SecondaryMap(num, banner)
	}
	for _, num := range b {
		str2 += SecondaryMap(num, banner)
	}
	for _, num := range c {
		str3 += SecondaryMap(num, banner)
	}
	for _, num := range d {
		str4 += SecondaryMap(num, banner)
	}
	for _, num := range e {
		str5 += SecondaryMap(num, banner)
	}
	for _, num := range f {
		str6 += SecondaryMap(num, banner)
	}
	for _, num := range g {
		str7 += SecondaryMap(num, banner)
	}
	for _, num := range h {
		str8 += SecondaryMap(num, banner)
	}
	return " " + str1 + "\n", str2 + "\n", str3 + "\n", str4 + "\n", str5 + "\n", str6 + "\n", str7 + "\n", str8
}
