package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	rgx, in := parseInput(input)
	fmt.Println(diffLengthCheck(rgx, in))
}

func getInput() ([]byte, error) {
	scanner := bufio.NewReader(os.Stdin)
	b, err := scanner.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("input: %v", err)
	}
	return b, nil
}

func parseInput(b []byte) ([]byte, []byte) {
	trimmed := bytes.TrimSpace(b)
	nb := bytes.Split(trimmed, []byte("|"))
	regex := nb[0]
	input := nb[1]
	return regex, input
}

func isMatching(r, i []byte) bool {
	regex := string(r)
	input := string(i)

	var match bool
	switch {
	case regex == "" && input == "",
		regex == "",
		regex == ".",
		regex == input:
		match = true
	case input == "":
		match = false
	default:
		match = false
	}
	return match
}

func matchingCharacters(r, i []byte) bool {
	if len(r) > 0 && len(i) <= 0 {
		return false
	}
	if len(r) > 0 && len(i) > 0 {
		if !isMatching(r[:1], i[:1]) {
			return false
		}
		return matchingCharacters(r[1:], i[1:])
	}
	return true
}

func diffLengthCheck(r, i []byte) bool {
	//fmt.Printf("Input: '%s|%s'\tOutput: %t\n", r, i, matchingCharacters(r, i))
	if matchingCharacters(r, i) {
		return true
	}
	if len(i) == 0 {
		return false
	}
	return diffLengthCheck(r, i[1:])
}
