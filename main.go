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
	err := initMain()
	parseFlags()

	infoStruct := []structs.Line{
		{Name: "OS", Value: utils.GetOsVersion()},
		{Name: "CPU", Value: utils.GetCpu()},
		{Name: "Uptime", Value: utils.GetUptime()},
		{Name: "Terminal", Value: utils.GetTerminal()},
		{Name: "Shell", Value: utils.GetShell()},
		{Name: "Root", Value: utils.GetDiskUsage("/")},
	}

	username := color.GreenString(utils.GetUsername()) + "@" + color.GreenString(utils.GetHostname())
	usernameLength := len([]rune(utils.StripAnsi(username)))
	separator := color.RGB(128, 128, 128).Sprint(strings.Repeat("-", usernameLength))

	info := formatLines(infoStruct)
	info = "\n" + username + "\n" + separator + "\n" + info

	lines := strings.Split(info, "\n")

	icon, iconColor := utils.GetIcon(utils.GetOS())
	iconLines := strings.Split(icon, "\n")

	lineCount := math.MaxInt(len(lines), len(iconLines))
	output := ""

	termWidth, _, err := term.GetSize(0)

	if err != nil {
		termWidth = 999
	}

	for index := range lineCount {
		line := ""

		if len(iconLines) > index {
			val := []rune(iconLines[index] + " ")
			valRaw := []rune(utils.StripAnsi(iconLines[index] + " "))

			if len(valRaw) > termWidth {
				val = val[:termWidth]
			}

			line += color.RGB(
				iconColor.R,
				iconColor.G,
				iconColor.B,
			).Sprint(string(val))
		}

		if len(lines) > index {
			val := lines[index]
			line += color.WhiteString(val)

			lineRaw := []rune(utils.StripAnsi(line))

			if len(lineRaw) > termWidth {
				line = utils.TrimAnsiString(line, termWidth)
			}
		}

		line += "\n"

		output += line
	}

	fmt.Print(output)
}

func initMain() error {
	term := utils.GetTerminal()

	if !strings.Contains(term, "color") {
		color.NoColor = true
	}

	return nil
}

func parseFlags() {
	noColour := flag.Bool("no-color", false, "Disable color output")

	flag.Parse()

	if *noColour {
		color.NoColor = true
	}
}

func formatLines(lines []structs.Line) string {
	output := ""
	longestNameLength := 0

	for _, line := range lines {
		if len(line.Name) > longestNameLength {
			longestNameLength = len(line.Name)
		}
	}

	for _, line := range lines {
		output += color.New(color.Bold, color.FgHiBlack).Sprintf("%s:", line.Name)
		spaces := longestNameLength - len(line.Name) + 1

		for spaces > 0 {
			output += " "
			spaces -= 1
		}

		output += line.Value + "\n"
	}

	return output
}
