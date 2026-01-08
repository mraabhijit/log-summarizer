package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func hard() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		os.Exit(1)
	}

	defer file.Close()

	searchTerms := os.Args[2:]

	lineCount := 0
	maxPatternLen := 0
	targets := make([][]byte, len(searchTerms))
	targetCounts := make([]int, len(searchTerms))
	onNewLogicalLine := true
	patternFoundOnLine := make([]bool, len(searchTerms))

	for i, term := range searchTerms {
		targets[i] = []byte(term)
		targetCounts[i] = 0
		if len(term) > maxPatternLen {
			maxPatternLen = len(term) // find the maximum length of the term to search for
		}
	}

	carry := make([]byte, 0, maxPatternLen) // start at 0 length slice with capacity = maxPatternlen
	reader := bufio.NewReader(file)

	for {
		chunk, err := reader.ReadBytes('\n')
		hasNewLine := bytes.HasSuffix(chunk, []byte("\n"))

		if len(chunk) > 0 {
			if onNewLogicalLine {
				lineCount++
				// since line already explored, set the NewLogicalLine to false
				onNewLogicalLine = false
			}

			// search for the pattern on the logical line if pattern not already found on the line
			// chunk... unpacks all bytes of chunk and appends to the existing bytes of carry
			searchSpace := append(carry, chunk...)
			for i, target := range targets {
				if !patternFoundOnLine[i] {
					if bytes.Contains(searchSpace, target) {
						targetCounts[i]++
						patternFoundOnLine[i] = true
					}
				}
			}

			if len(searchSpace) >= maxPatternLen {
				// if pattern is "ABCDE", and the chunk is "...ABCD",
				// then "E" should be in the next chunk, so we use maxPatternLen-1 to carry forward
				// suppose the next chunk is "EBCFEO", then
				// after appending to carry, SearchSpace becomes "ABCDEBCFEO", which
				// helps in finding the pattern across chunks
				carry = append(carry[:0], searchSpace[len(searchSpace)-maxPatternLen+1:]...)
			} else {
				carry = append(carry[:0], searchSpace...)
			}

		}

		if err != nil {
			if err == io.EOF {
				break
			}
			os.Exit(1)
		}

		// start of a new line should make the onNewLogicalLine flag true
		// also reset all target findings to false as they have not been found yet
		if hasNewLine {
			onNewLogicalLine = true
			carry = carry[:0]
			for i := range patternFoundOnLine {
				patternFoundOnLine[i] = false
			}
		}
	}
	print(lineCount)
	for i := range searchTerms {
		print(" ", targetCounts[i])
	}
}
