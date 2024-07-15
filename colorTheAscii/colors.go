package colortheascii

func GetColor(color string) (code int) {
	switch color {
	case "red":
		return 31
	case "#ff0000":
		return 31
	case "green":
		return 32
	case "#00ff00":
		return 32
	case "yellow":
		return 33
	case "#ffff00":
		return 33
	case "blue":
		return 34
	case "#0000ff":
		return 34

	case "magenta":
		return 35
	case "#ff00ff":
		return 35
	case "cyan":
		return 36
	case "#00ffff":
		return 36
	case "orange":
		return 214
	case "#ffa500":
		return 214
	case "black":
		return 30
	case "#000000":
		return 30
	case "purple":
		return 129
	case "#800080":
		return 129
	case "pink":
		return 218
	case "#ffc0cb":
		return 218
	case "#ffffff":
		return 0
	default:
		return 0
	}
}

/*

 */
