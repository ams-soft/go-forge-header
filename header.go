package header

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ams-soft/tic"
)

const TotalColumns = 80

func printLine() string {
	return strings.Repeat("-", TotalColumns)
}

func printSpace() string {
	return strings.Repeat(" ", TotalColumns)
}

func printSymbol() string {
	red := tic.Style().Bold().Fg(tic.ColorLightRed).Sprint("⫽")
	orange := tic.Style().Bold().Fg(tic.ColorOrange).Sprint("⫽")
	yellow := tic.Style().Bold().Fg(tic.ColorYellow).Sprint("⫽")
	green := tic.Style().Bold().Fg(tic.ColorLightGreen).Sprint("⫽")
	blue := tic.Style().Bold().Fg(tic.ColorLightBlue).Sprint("⫽")

	return red + orange + yellow + green + blue
}

func isColor() bool {
	b, err := strconv.ParseBool(os.Getenv("GO_FORGE_HEADER_NO_COLOR"))
	if err != nil {
		return true
	}
	return !b
}

func AmsLogo() string {
	if isColor() {
		return amsLogoColor()
	}
	return "AMS SOFT"
}

func amsLogoColor() string {
	a := tic.Style().Bold().Italic().Sprint("A")
	ms := tic.Style().Bold().Sprint("MS SOFT")

	str := printSymbol() + a + ms
	return str
}

func Buffer(vendor *string, name, version, env string, utc bool) []string {
	now := time.Now()
	if utc {
		now = now.UTC()
	}

	vendorStr := AmsLogo()
	if vendor != nil {
		vendorStr = *vendor
	}

	buffer := []string{}

	buffer = append(buffer, printLine())
	buffer = append(buffer, fmt.Sprintf("%s - %s", vendorStr, strings.ToUpper(name)))
	buffer = append(buffer, printLine())
	buffer = append(buffer, printSpace())
	buffer = append(buffer, printLine())
	buffer = append(buffer, "SYSTEM INFORMATION")
	buffer = append(buffer, printLine())
	buffer = append(buffer, fmt.Sprintf("VERSION:          %s", strings.ToUpper(version)))
	buffer = append(buffer, fmt.Sprintf("MODE:             %s", strings.ToUpper(env)))
	buffer = append(buffer, fmt.Sprintf("DATE:             %s", now.Format(time.RFC3339)))
	buffer = append(buffer, printSpace())

	return buffer
}

func Print(vendor *string, name, version, env string, utc bool) {
	buffer := Buffer(vendor, name, version, env, utc)
	for _, line := range buffer {
		fmt.Println(line)
	}
}
