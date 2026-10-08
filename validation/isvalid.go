package validation

import (
	"fmt"
	"os"
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
		return
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(usage)
		return
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	input := strings.Split(content, "\n")

	// fmt.Println(input[0])

	ants, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println("ERROR: invalid data format, invalid number of ants")
		return
	}
	if ants <= 0 {
		fmt.Println("ERROR: invalid data format, number of ants should be greater than or equal to 1")
		return
	}

	var cleaned []string

	for _, line := range input {
		if strings.HasPrefix(line, "#") {
			if line != "##start" && line != "##end" {
				continue
			}
		}

		cleaned = append(cleaned, line)
	}

	input = cleaned

	foundStart := false
	foundEnd := false

	for _, v := range input {
		if v == "##start" {
			if foundStart {
				fmt.Println("ERROR: invalid data format, only 1 start command allowed")
				return
			}
			foundStart = true
		}

		if v == "##end" {
			if foundEnd {
				fmt.Println("ERROR: invalid data format, only 1 end command allowed")
				return
			}
			foundEnd = true
		}
	}

	if !foundStart || !foundEnd {
		fmt.Println("ERROR: invalid data format")
		return
	}

	for i, v := range input {
		if v == "##start" || v == "##end" {
			if i+1 >= len(input) {
				fmt.Println("ERROR: invalid data format, missing room")
				return
			}

			fields := strings.Fields(input[i+1])

			if len(fields) != 3 {
				fmt.Println("ERROR: invalid data format, invalid room")
				return
			}
		}
	}

	fmt.Println(input)
}
