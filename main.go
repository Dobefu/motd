package main

import (
	"fmt"
	"motd/math"
	"motd/structs"
	"motd/utils"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func main() {
	info_struct := []structs.Line{
		{Name: "OS", Value: utils.GetOsVersion()},
	}

	info := formatLines(info_struct)

	info = figure.NewFigure(utils.GetUsername(), "doom", false).String() + "\n" + info

	lines := strings.Split(info, "\n")

	icon, icon_color := utils.GetIcon(utils.GetOS())
	icon_lines := strings.Split(icon, "\n")

	line_count := math.MaxInt(len(lines), len(icon_lines))
	output := ""

	for index := range line_count {
		if len(icon_lines) > index {
			output += color.RGB(
				icon_color.R,
				icon_color.G,
				icon_color.B,
			).Sprint(icon_lines[index]) + " "
		}

		if len(lines) > index {
			output += color.WhiteString(lines[index])
		}

		output += "\n"
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
