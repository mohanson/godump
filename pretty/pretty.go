package pretty

import (
	"fmt"
	"strings"
)

// PrintTable easily draw tables in terminal/console applications from a list of lists of strings.
func PrintTable(data [][]string) {
	size := make([]int, len(data[0]))
	for _, r := range data {
		for j, c := range r {
			size[j] = max(size[j], len(c))
		}
	}
	line := make([]string, len(data[0]))
	for j, c := range data[0] {
		l := size[j]
		line[j] = strings.Repeat(" ", l-len(c)) + c
	}
	fmt.Println(strings.Join(line, " "))
	for i, c := range size {
		line[i] = strings.Repeat("-", c)
	}
	fmt.Println(strings.Join(line, "-"))
	for _, r := range data[1:] {
		for j, c := range r {
			l := size[j]
			line[j] = strings.Repeat(" ", l-len(c)) + c
		}
		fmt.Println(strings.Join(line, " "))
	}
}
