package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	searchTarget := flag.String("word", "", "flag 1")
	target := flag.String("target", "", "flag 2")
	flag.Parse()

	execGrep(*searchTarget, *target)
}

func execGrep(text string, target string) {
	res, _ := exec.Command("grep", "-r", text, target).Output()
	result := formatGrepResult(res)

	fmt.Printf("Search text : %v\n", text)
	output(result)

	if len(result) > 0 {
		for _, v := range result {
			newText := strings.Replace(v, target + "/", "", 1)
			execGrep(newText, target)
		}
	} else {
		fmt.Println("No results.")
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
