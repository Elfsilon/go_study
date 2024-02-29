package visualizer

import (
	"strings"
)

const (
	Reset string = "\x1b[0m"

	Black        string = "\x1b[30m"
	DarkGray     string = "\x1b[30;1m"
	Red          string = "\x1b[31m"
	VividRed     string = "\x1b[31;1m"
	Green        string = "\x1b[32m"
	VividGreen   string = "\x1b[32;1m"
	Yellow       string = "\x1b[33m"
	VividYellow  string = "\x1b[33;1m"
	Blue         string = "\x1b[34m"
	VividBlue    string = "\x1b[34;1m"
	Magenta      string = "\x1b[35m"
	VividMagenta string = "\x1b[35;1m"
	Cyan         string = "\x1b[36m"
	VividCyan    string = "\x1b[36;1m"
	Gray         string = "\x1b[37m"
	White        string = "\x1b[37;1m"

	BlackBG   string = "\x1b[40m"
	RedBG     string = "\x1b[41m"
	GreenBG   string = "\x1b[42m"
	YellowBG  string = "\x1b[43m"
	BlueBG    string = "\x1b[44m"
	MagentaBG string = "\x1b[45m"
	CyanBG    string = "\x1b[46m"
	GrayBG    string = "\x1b[47m"

	Bold      string = "\x1b[1m"
	Italic    string = "\x1b[3m"
	Underline string = "\x1b[4m"
	Blink     string = "\x1b[5m"
	Inverse   string = "\x1b[7m"
)

func ApplyColor(text string, color string) string {
	var b strings.Builder

	b.WriteString(color)
	b.WriteString(Bold)
	b.WriteString(text)
	b.WriteString(Reset)

	return b.String()
}
