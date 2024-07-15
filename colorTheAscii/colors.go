package colortheascii

func GetColor(color string) (code int) {
	switch color {
	case "red":
		return 31
	case "#ff0000":
		return 31
	case "rgb(255, 0, 0)":
		return 31
	case "hsl(0, 100%, 50%)":
		return 31
	case "green":
		return 32
	case "#00ff00":
		return 32
	case "rgb(0, 255, 0)":
		return 32
	case "hsl(120, 100%, 50%)":
		return 32
	case "yellow":
		return 33
	case "#ffff00":
		return 33
	case "rgb(255, 255, 0)":
		return 33
	case "hsl(60, 100%, 50%)":
		return 33
	case "blue":
		return 34
	case "#0000ff":
		return 34
	case "rgb(0, 0, 255)":
		return 34
	case "hsl(240, 100%, 50%)":
		return 34
	case "magenta":
		return 35
	case "#ff00ff":
		return 35
	case "rgb(255, 0, 255)":
		return 35
	case "hsl(300, 100%, 50%)":
		return 35
	case "cyan":
		return 36
	case "#00ffff":
		return 36
	case "rgb(0, 255, 255)":
		return 36
	case "hsl(180, 100%, 50%)":
		return 36
	case "orange":
		return 214
	case "#ffa500":
		return 214
	case "rgb(255, 165, 0)":
		return 214
	case "hsl(30, 100%, 50%)":
		return 214
	case "black":
		return 30
	case "#000000":
		return 30
	case "rgb(0, 0, 0)":
		return 30
	case "hsl(0, 0%, 0%)":
		return 30
	case "purple":
		return 129
	case "#800080":
		return 129
	case "rgb(128, 0, 128)":
		return 129
	case "hsl(300, 100%, 25%)":
		return 129
	case "pink":
		return 218
	case "#ffc0cb":
		return 218
	case "rgb(255, 192, 203)":
		return 218
	case "hsl(350, 100%, 88%)":
		return 218
	default:
		return 0
	}
}
