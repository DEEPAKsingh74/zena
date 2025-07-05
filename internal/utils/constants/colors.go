package constants

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorOrange = "\033[33m"
	ColorYellow = "\033[93m"
	ColorWhite  = "\033[97m"
)

var ColorMap = map[string]string{
	"red":    ColorRed,
	"green":  ColorGreen,
	"blue":   ColorBlue,
	"orange": ColorOrange,
	"yellow": ColorYellow,
	"white":  ColorWhite,
}
