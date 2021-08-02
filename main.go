package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	searchTarget := flag.String("word", "", "search target word.")
	targetDir := flag.String("targetDir", "", "target directory.")
	removePath := flag.String("removePath", "", "remove path prefix.")
	flag.Parse()

	execGrep(*searchTarget, *targetDir, *removePath)
}

func execGrep(text string, target string, removePath string) {
	fmt.Printf("Search text : %v\n", text)

	res, _ := exec.Command("grep", "-r", "--exclude-dir", ".git", text, target).Output()
	result := formatGrepResult(res)

	output(result)

	if len(result) > 0 {
		fmt.Printf("\n")
		for _, v := range result {
			newText := strings.Replace(v, removePath, "", 1)
			execGrep(newText, target, removePath)
		}
	} else {
		fmt.Println("No results.")
		fmt.Printf("\n")
	}

}

func formatGrepResult(grepResult []byte) []string {
	resultList := strings.Split(string(grepResult), "\n")

	m := make(map[string]struct{})
	newList := make([]string, 0)

	for _, v := range resultList {
		if len(v) > 0 {
			path := strings.Split(v, ":")[0]
			if _, ok := m[path]; !ok {
				m[path] = struct{}{}
				newList = append(newList, path)
			}
		}
	}

	return newList
}

func output(line []string) {
	for _, v := range line {
		fmt.Println(v)
	}
}
