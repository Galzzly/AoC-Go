package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadLines(f string) (line []string, err error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	line = make([]string, 0, len(lines))

	for l := range lines {
		if len(lines[l]) == 0 {
			continue
		}
	}
	return line, nil
}

func main() {
	f := os.Args[1]
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	for l := range lines {
		res := strings.Split(l, " ")
		for s := range res {
			fmt.Printf("%s\n", s)
		}
	}
	fmt.Println("vim-go")
}
