package colortheascii

func GetColor(color string) (code int) {
	switch color {
	case "red":
		return 31
	case "green":
		return 32
	case "yellow":
		return 33
	case "blue":
		return 34
	case "magenta":
		return 35
	case "cyan":
		return 36
	case "black":
		return 30
	case "white":
		return 37
	case "redbackground":
		return 41
	case "greenbackground":
		return 42
	case "yellowbackground":
		return 43
	case "bluebackground":
		return 44
	case "magentabackground":
		return 45
	case "cyanbackground":
		return 46
	case "blackbackground":
		return 40
	case "whitebackground":
		return 47
	default:
		return 0
	}
}
