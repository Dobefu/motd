package main

import (
	"flag"
	"fmt"
	"motd/math"
	"motd/structs"
	"motd/utils"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func main() {
	parseFlags()

	info_struct := []structs.Line{
		{Name: "OS", Value: utils.GetOsVersion()},
		{Name: "CPU", Value: utils.GetCpu()},
		{Name: "Uptime", Value: utils.GetUptime()},
		{Name: "Terminal", Value: utils.GetTerminal()},
		{Name: "Shell", Value: utils.GetShell()},
		{Name: "Root", Value: utils.GetDiskUsage("/")},
	}

	username := color.GreenString(utils.GetUsername()) + "@" + color.GreenString(utils.GetHostname())
	username_length := len(utils.GetUsername()) + len(utils.GetHostname()) + 1
	separator := color.RGB(128, 128, 128).Sprint(strings.Repeat("-", username_length))

	info := formatLines(info_struct)
	info = "\n" + username + "\n" + separator + "\n" + info

	lines := strings.Split(info, "\n")

	icon, icon_color := utils.GetIcon(utils.GetOS())
	icon_lines := strings.Split(icon, "\n")

	line_count := math.MaxInt(len(lines), len(icon_lines))
	output := ""

	term_width, _, err := term.GetSize(0)

	if err != nil {
		term_width = 999
	}

	for index := range line_count {
		line_raw := ""
		line := ""

		if len(icon_lines) > index {
			val := []rune(icon_lines[index] + " ")

			if len(val) > term_width {
				val = val[:term_width]
			}

			line_raw += string(val)
			line += color.RGB(
				icon_color.R,
				icon_color.G,
				icon_color.B,
			).Sprint(string(val))
		}

		if len(lines) > index {
			val := lines[index]

			if len(val) > term_width-len([]rune(line_raw)) {
				limit := term_width - len([]rune(line_raw))

				if limit < 0 {
					limit = 0
				}

				val = val[:limit]
			}

			line_raw += val

			val_parts := strings.Split(val, ": ")

			if len(val_parts) > 1 {
				val = fmt.Sprintf(
					"%s %s",
					color.New(color.Bold, color.FgHiBlack).Sprintf("%s:", val_parts[0]),
					val_parts[1],
				)
			}

			line += color.WhiteString(val)

		}

		line += "\n"

		output += line
	}

	fmt.Print(output)
}

func parseFlags() {
	no_colour := flag.Bool("no-color", false, "Disable color output")

	flag.Parse()

	if *no_colour {
		color.NoColor = true
	}
}

func formatLines(lines []structs.Line) string {
	output := ""
	longest_name_length := 0

	for _, line := range lines {
		if len(line.Name) > longest_name_length {
			longest_name_length = len(line.Name)
		}
	}

	for _, line := range lines {
		output += line.Name + ":"
		spaces := longest_name_length - len(line.Name) + 1

		for spaces > 0 {
			output += " "
			spaces -= 1
		}

		output += line.Value + "\n"
	}

	return output
}
