package helpers

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// PrintInfo print Info
func PrintInfo() {
	f := figure.NewColorFigure("Clipboard-Sync", "big", "red", true)
	figletStr := f.String()
	fmt.Println(figletStr)
	fmt.Println()

}
