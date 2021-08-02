package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

type Result struct {
	globalMap  map[string]struct{}
	globalList []string
	depth      int
}

func main() {
	searchTarget := flag.String("word", "", "search target word.")
	targetDir := flag.String("targetDir", "", "target directory.")
	removePath := flag.String("removePath", "", "remove path prefix.")
	flag.Parse()

	// Holds the Map and Slice as structs for the overall result.
	result := Result{
		globalMap:  make(map[string]struct{}),
		globalList: make([]string, 0),
		depth:      0,
	}

	execGrep(*searchTarget, *targetDir, *removePath, &result)

	fmt.Printf("---- Result ---\n")
	fmt.Printf("Target Files : %v\n", len(result.globalList))
	output(result.globalList, 0)
}

func execGrep(text string, target string, removePath string, globalResult *Result) {


	fmt.Printf("depth : %v\n", globalResult.depth)
	fmt.Printf("Search text : %v\n", text)

	// Execute grep.
	res, _ := exec.Command("grep", "-r", "--exclude-dir", ".git", text, target).Output()
	// Format grep results.
	result := formatGrepResult(res, globalResult)

	output(result, globalResult.depth)

	if len(result) > 0 {
		fmt.Printf("\n")
		globalResult.depth++
		for _, v := range result {
			newText := strings.Replace(v, removePath, "", 1)
			execGrep(newText, target, removePath, globalResult)
		}
		globalResult.depth--
	} else {
		fmt.Println("No results.")
		fmt.Printf("\n")
	}
}

// Split grep results to eliminate duplicate results.
func formatGrepResult(grepResult []byte, globalResult *Result) []string {
	resultList := strings.Split(string(grepResult), "\n")

	m := make(map[string]struct{})
	newList := make([]string, 0)

	for _, v := range resultList {
		if len(v) > 0 {
			path := strings.Split(v, ":")[0]
			// Duplicate check for current search terms.
			if _, ok := m[path]; !ok {
				m[path] = struct{}{}
				newList = append(newList, path)
			}
			// Duplicate checking for the entire result.
			if _, ok := globalResult.globalMap[path]; !ok {
				globalResult.globalMap[path] = struct{}{}
				globalResult.globalList = append(globalResult.globalList, path)
			}
		}
	}

	return newList
}

func output(line []string, depth int) {
	var printPrefix = ""
	if depth > 0 {
		printPrefix = "  |-="
	}

	for _, v := range line {
		fmt.Printf("%v%v\n", printPrefix, v)
	}
}
