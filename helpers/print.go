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

func MaskString(s string, mask rune, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(s) || end < 0 {
		end = len(s)
	}
	if start > end {
		start, end = end, start
	}
	return s[:start] + maskString(s[start:end], mask) + s[end:]
}

func maskString(s string, mask rune) string {
	r := make([]rune, len(s))
	for i := range r {
		r[i] = mask
	}
	return string(r)
}
