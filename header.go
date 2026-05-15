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

func Print(vendor *string, name, version, env string, utc bool) {
	now := time.Now()
	if utc {
		now = now.UTC()
	}

	vendorStr := AmsLogo()
	if vendor != nil {
		vendorStr = *vendor
	}
	fmt.Println(printLine())
	fmt.Printf("%s - %s\n", vendorStr, strings.ToUpper(name))
	fmt.Println(printLine())
	fmt.Println(printSpace())
	fmt.Println(printLine())
	fmt.Println("SYSTEM INFORMATION")
	fmt.Println(printLine())
	fmt.Println("VERSION:          ", strings.ToUpper(version))
	fmt.Println("MODE:             ", strings.ToUpper(env))
	fmt.Println("DATE:             ", now.Format(time.RFC3339))
	fmt.Println(printSpace())
}
