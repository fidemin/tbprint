package tbprint

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxWidths(input [][]string) map[int]int {
	widths := make(map[int]int)
	for _, row := range input {
		for j, col := range row {
			if oldLen, ok := widths[j]; !ok {
				widths[j] = len(col)
			} else {
				widths[j] = max(oldLen, len(col))
			}
		}
	}
	return widths
}

func Fprint(w io.Writer, input [][]string) {
	// maxWidths is max width of nth column
	widths := maxWidths(input)
	for _, row := range input {
		for j, col := range row {
			restLen := widths[j] - len(col)
			fmt.Fprintf(w, "%s%s", col, strings.Repeat(" ", restLen+2))
		}
		fmt.Fprint(w, "\n")
	}
}

func Print(input [][]string) {
	Fprint(os.Stdout, input)
}
