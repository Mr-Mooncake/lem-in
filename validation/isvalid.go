package validation

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const usage = "Usage: go run . <filename.txt>"

func IsValid() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	filename := args[1]

	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println(usage)
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(usage)
	}
	input := strings.Split(string(data), "\r\n")

	// fmt.Println(input[0])

	_, err = strconv.Atoi(input[0])
	if err != nil {
		fmt.Println("Error: Numbers of ants should be an integer")
		return
	}

	found := make(map[string]int)

	for i := 0; i < len(input); {
		if strings.Count(input[i], "#") == 1 {
			input = slices.Delete(input, i, i+1)
		} else {
			i++
		}
	}

	for i, v := range input {
		if v == "##start" {
			found["##start"] = i
		} else if v == "##end" {
			found["##end"] = i
		} else if strings.Contains(v, "#") {
			fmt.Println("Invalid format:", v)
			return
		}
	}

	_, foundStart := found["##start"]
	_, foundEnd := found["##end"]

	if !foundStart || !foundEnd {
		fmt.Println("Both ##start and ##end must exist")
		return
	}

	if found["##start"] >= found["##end"] {
		fmt.Println("##start must come before ##end")
		return
	}
	fmt.Println(input)
}
