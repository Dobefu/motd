package main

import (
	"fmt"
	"motd/structs"
	"motd/utils"
)

func main() {
	lines := []structs.Line{
		{Name: "OS", Value: utils.GetOsVersion()},
	}

	formatLines(lines)

	for _, line := range lines {
		fmt.Print(line.Name)
		fmt.Println(line.Value)
	}

	icon := utils.GetIcon(utils.GetOS())
	fmt.Printf(icon)
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
