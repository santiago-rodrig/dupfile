// Dup2 prints the count and text of lines that appear more than once
// int the input. It reads from stdin or from a list of named files.
// This version also prints the name of the files where the duplication
// happens.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type fileCounter struct {
	fname  string
	counts map[string]int
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var fileCounts []*fileCounter
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			fCounter := &fileCounter{fname: arg, counts: make(map[string]int)}
			fileCounts = append(fileCounts, fCounter)
			countLines(f, fCounter.counts)
			f.Close()
		}
	}
	for _, fCounter := range fileCounts {
		fmt.Printf("%s:\n", fCounter.fname)
		hasDuplicates := false
		for line, n := range fCounter.counts {
			if n > 1 {
				hasDuplicates = true
				fmt.Printf("\t%d: %s\n", n, line)
			}
		}
		if !hasDuplicates {
			fmt.Println("\tNO DUPLICATE LINES")
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
