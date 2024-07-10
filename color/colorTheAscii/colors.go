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
	default:
		return 0
	}
}
