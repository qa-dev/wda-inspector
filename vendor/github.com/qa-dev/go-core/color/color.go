package color

import "fmt"

type TypeColor string

const (
	Black        TypeColor = "\x1b[30m"
	Red          TypeColor = "\x1b[31m"
	Green        TypeColor = "\x1b[32m"
	Yellow       TypeColor = "\x1b[33m"
	Blue         TypeColor = "\x1b[34m"
	Magenta      TypeColor = "\x1b[35m"
	Cyan         TypeColor = "\x1b[36m"
	White        TypeColor = "\x1b[37m"
	Default      TypeColor = "\x1b[39m"
	LightGray    TypeColor = "\x1b[90m"
	LightRed     TypeColor = "\x1b[91m"
	LightGreen   TypeColor = "\x1b[92m"
	LightYellow  TypeColor = "\x1b[93m"
	LightBlue    TypeColor = "\x1b[94m"
	LightMagenta TypeColor = "\x1b[95m"
	LightCyan    TypeColor = "\x1b[96m"
	LightWhite   TypeColor = "\x1b[97m"
)

var def TypeColor = Default

func Color(c TypeColor, s string) string {
	return fmt.Sprintf("%s%s%s", c, s, DefColor())
}

func ColorN(c TypeColor, s string) string {
	return Color(c, s) + "\n"
}

func SetDefColor(c TypeColor) {
	def = c
}

func DefColor() TypeColor {
	return def
}
