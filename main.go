package main

import (
	"fmt"
	"motd/math"
	"motd/structs"
	"motd/utils"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"golang.org/x/term"
)

func main() {
	info_struct := []structs.Line{
		{Name: "OS", Value: utils.GetOsVersion()},
		{Name: "CPU", Value: utils.GetCpu()},
		{Name: "Uptime", Value: utils.GetUptime()},
		{Name: "Terminal", Value: utils.GetTerminal()},
		{Name: "Root", Value: utils.GetDiskUsage("/")},
	}

	info := formatLines(info_struct)
	info = figure.NewFigure(utils.GetUsername(), "doom", false).String() + "\n" + info

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
					color.New(color.Bold).Sprintf("%s:", val_parts[0]),
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
