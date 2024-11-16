package main

import (
	"fmt"
	"motd/math"
	"motd/structs"
	"motd/utils"
	"strings"

	"github.com/fatih/color"
)

func main() {
	lines := []structs.Line{
		{Name: "OS", Value: utils.GetOsVersion()},
	}

	formatLines(lines)

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
			).Sprint(icon_lines[index])
		}

		if len(lines) > index {
			output += color.WhiteString(lines[index].Name + lines[index].Value)
		}

		output += "\n"
	}

	fmt.Print(output)
}

func formatLines(lines []structs.Line) {
	longest_name_length := 0

	for _, line := range lines {
		if len(line.Name) > longest_name_length {
			longest_name_length = len(line.Name)
		}
	}

	for index, line := range lines {
		spaces := longest_name_length - len(line.Name) + 1
		lines[index].Name += ":"

		for spaces > 0 {
			lines[index].Name += " "
			spaces -= 1
		}
	}
}
