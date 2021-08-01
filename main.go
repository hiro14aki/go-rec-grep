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
	res, _ := exec.Command("grep", "-r", *searchTarget, *target).Output()

	resultList := strings.Split(string(res), "\n")
	fmt.Println(len(resultList))

	m := make(map[string]struct{})
	newList := make([]string, 0)

	for _, v := range resultList {
		path := strings.Split(v, ":")[0]

		if _, ok := m[path]; !ok {
			m[path] = struct{}{}
			newList = append(newList, path)
		}
	}

	fmt.Println(len(newList))
}

