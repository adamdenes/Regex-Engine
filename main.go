package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	rgx, in := parseInput(input)
	fmt.Println(isMatching(rgx, in))
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
	//fmt.Printf("nb[0]=%v, nb[1]=%v\n", nb[0], nb[1])
	//fmt.Printf("str nb[0]=%v, str nb[1]=%v\n", string(nb[0]), string(nb[1]))
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
		strings.Compare(regex, input) == 0:
		match = true
	case input == "":
		match = false
	default:
		match = false
	}
	return match
}
