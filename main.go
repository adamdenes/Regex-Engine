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
	before, after, _ := bytes.Cut(trimmed, []byte("|"))
	return before, after
}

func isMatching(r, i []byte) bool {
	regex := string(r[:1])
	input := string(i[:1])
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
		if len(r) == 1 && string(r[len(r)-1:]) == "$" {
			return true
		}
		return false
	}

	if len(r) > 0 && len(i) > 0 {
		if string(r[:1]) == "\\" {
			r = r[1:]
		}
		if !isMatching(r, i) {
			return false
		}
		return matchingCharacters(r[1:], i[1:])
	}
	return true
}

func diffLengthCheck(r, i []byte) bool {
	fmt.Printf("Input:\t'%s|%s'\tOutput: %t\n", r, i, matchingCharacters(r, i))
	if len(r) == 0 {
		return true
	}

	ch, idx := useMetaChar(r)

	switch string(ch) {
	case "?":
		matchZeroOrOnce(idx, &r, &i)
	case "*":
		matchZeroOrMore(idx, &r, &i)
	case "+":
		matchOnceOrMore(idx, &r, &i)
	}

	if strings.HasPrefix(string(r[:1][0]), "^") {
		return matchingCharacters(r[1:], i)
	}
	if matchingCharacters(r, i) {
		return true
	}
	if len(i) == 0 {
		return false
	}
	return diffLengthCheck(r, i[1:])
}

// provides the index of the meta character
func useMetaChar(b []byte) ([]byte, int) {
	var idx int
	var meta bool
	ch := make([]byte, 0, 2)

	for _, c := range b {
		switch c {
		case '?', '*', '+':
			meta = bytes.IndexByte(b, c) != -1
			if !meta {
				continue
			}
			idx = bytes.IndexByte(b, c)
			ch = append(ch, c)
		}
	}
	return ch, idx
}

func omit(idx int, r, i *[]byte) bool {
	var first, rest []byte
	// omit the meta char
	if (len(*r)-idx)-(len(*i)-idx) == 2 {
		first = (*r)[:idx-1]
		rest = (*r)[idx+1:]
		*r = append(first, rest...)
		return true
	}
	return false
}

// `?` matches the preceding character zero times or once
func matchZeroOrOnce(idx int, r, i *[]byte) {
	if string((*r)[idx-1]) == "\\" {
		return
	}

	if !omit(idx, r, i) && (*i)[idx-1] == (*r)[idx-1] ||
		string((*r)[idx-1]) == "." {
		first := (*r)[:idx]
		rest := (*r)[idx+1:]
		*r = append(first, rest...)
	}
}

// `*` matches the preceding character zero or more times
func matchZeroOrMore(idx int, r, i *[]byte) {
	if string((*r)[idx-1]) == "\\" {
		return
	}
	// something like '.*|aaa' would go out of length otherwise
	if len(*r) == 2 {
		first := (*r)[:idx]
		rest := (*r)[idx+1:]
		*r = append(first, rest...)
		return
	}

	if !omit(idx, r, i) && (*i)[idx-1] == (*r)[idx-1] ||
		string((*r)[idx-1]) == "." {

		*r = (*r)[idx-1:]
		*i = (*i)[idx-1:]
	}
}

// `+` matches the preceding character once or more times
func matchOnceOrMore(idx int, r, i *[]byte) {
	count := strings.Count(string(*i), string((*r)[idx-1]))
	// preceding char occurs once
	if count == 1 && (*i)[idx-1] == (*r)[idx-1] ||
		string((*r)[idx-1]) == "." {
		first := (*r)[:idx]
		rest := (*r)[idx+1:]
		*r = append(first, rest...)
	}

	if count > 1 {
		*r = (*r)[idx-1:]
		*i = (*i)[idx-1:]
	}
}
