package main

import (
	"fmt"
	"io/ioutil"
	//	"os"
	//	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getDets(instr string) (string, int) {
	splitOnSpace := strings.Split(instr, " ")
	operation := splitOnSpace[0]
	sign := 1

	n, _ := strconv.Atoi(splitOnSpace[1][1:])
	if strings.HasPrefix(splitOnSpace[1], "-") {
		sign = -1
	}
	operand := n * sign
	return operation, operand
}

func main() {
	f, err := ioutil.ReadFile("input.txt")
	check(err)
	var tochange []int

	lines := strings.Split(string(f), "\n")
	// Get the list of stuff that isn't acc
	tochange = make([]int, 0, len(lines))
	for l := range lines {
		s := strings.Split(lines[l], " ")
		if !(s[0] == "acc") {
			tochange = append(tochange, l)
		}
	}

	var acc = 0
	var atmpt = 0
	for tc := range tochange {
		acc = 0
		atmpt++
		fmt.Printf("attempt %d\n", atmpt)
		var ran []int
		i := 0
		for {
			if contains(ran, i) {
				break
			}
			ran = append(ran, i)
			op, nu := getDets(lines[i])
			switch op {
			case "acc":
				acc += nu
				i++
			case "jmp":
				if i == tc {
					i++
				} else {
					i += nu
				}
			case "nop":
				if i == tc {
					i += nu
				} else {
					i++
				}
			}
		}
		if i >= len(lines) {
			break
		}
		acc = 0
	}
	fmt.Printf("Found %d in %d attempts\n", acc, atmpt)
}
