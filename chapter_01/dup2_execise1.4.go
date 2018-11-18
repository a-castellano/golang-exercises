// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	filecounts := make(map[string]map[string]int)

	var fileappears string

	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(arg, f, counts, filecounts)
		f.Close()
	}

	for line, n := range counts {
		for filename, _ := range filecounts[line] {
			fileappears += filename + " "
		}
		if n > 1 {
			fmt.Printf("%d\t%s - %s\n", n, line, fileappears)
		}
		fileappears = ""
	}
}
func countLines(filename string, f *os.File, counts map[string]int, filecounts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if _, ok := filecounts[input.Text()]; !ok {
			filecounts[input.Text()] = map[string]int{}
		}
		filecounts[input.Text()][filename] = 1
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
