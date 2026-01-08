package main

import (
	"os"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}
	hard()
}

// 	fileName := os.Args[1]
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		os.Exit(1)
// 	}

// 	defer file.Close()

// 	searchTerms := os.Args[2:]

// 	lineCount := 0
// 	targets := make([][]byte, len(searchTerms))
// 	targetCounts := make([]int, len(searchTerms))

// 	for i, term := range searchTerms {
// 		targets[i] = []byte(term)
// 		targetCounts[i] = 0
// 	}

// 	scanner := bufio.NewScanner(file)

// 	const maxCapacity = 1024 * 1024 * 1048 // >1 GB
// 	buf := make([]byte, 64*1024)
// 	scanner.Buffer(buf, maxCapacity)

// 	for scanner.Scan() {
// 		lineCount++
// 		line := scanner.Bytes()

// 		for i, target := range targets {
// 			if bytes.Contains(line, target) {
// 				targetCounts[i]++
// 			}
// 		}
// 	}
// 	if scanner.Err() != nil {
// 		os.Exit(1)
// 	}

// 	print(lineCount)
// 	for i, term := range searchTerms {
// 		print(" ", term, " ", targetCounts[i])
// 	}
// }
